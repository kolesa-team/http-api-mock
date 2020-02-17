package vars

import (
	"github.com/kolesa-team/http-api-mock/definition"
	"github.com/kolesa-team/http-api-mock/persist"
	"github.com/kolesa-team/http-api-mock/utils"
	"regexp"
	"strings"
)

type DynamicVarsFiller struct {
	Engines     *persist.PersistEngineBag
	RegexHelper utils.RegexHelper
}

func (dvf DynamicVarsFiller) Fill(m *definition.Mock, input string, multipleMatch bool) string {
	r := regexp.MustCompile(`\{\{\s*dynamic\.([^{]+?)\s*\}\}`)
	tries := 0
	// we are making several passes while we have matching regex,
	// this is useful for cases when we have nested vars like
	// {{ dynamic.Calc({{ request.body.term1 }} + {{ dynamic.Calc({{ request.body.term2 }}) }}) }}
	for tries <= 3 && r.MatchString(input) {
		input = dvf.Process(r, m, input)
		tries++
	}

	return input
}

func (dvf DynamicVarsFiller) Process(r *regexp.Regexp, m *definition.Mock, input string) string {
	return r.ReplaceAllStringFunc(input, func(raw string) string {
		found := false
		s := ""
		tag := strings.Trim(raw[2:len(raw)-2], " ")
		if i := strings.Index(tag, "dynamic.Calc"); i == 0 {
			s, found = dvf.callCalc(m, tag[len("dynamic.Calc"):])
		}

		if !found {
			return raw
		}
		return s
	})
}

func (dvf DynamicVarsFiller) callCalc(m *definition.Mock, parameters string) (string, bool) {
	regexPattern := `\(\s*(?:'|")?(?P<expression>.+?)(?:'|")?\s*\)`

	helper := utils.RegexHelper{}

	expression, found := helper.GetStringPart(parameters, regexPattern, "expression")
	if !found {
		return "", false
	}

	if result, err := utils.CalcExpression(expression); err == nil {
		return result, true
	} else {
		return "", false
	}
}
