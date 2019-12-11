package vars

import "github.com/kolesa-team/http-api-mock/definition"

type Filler interface {
	Fill(m *definition.Mock, input string, multipleMatch bool) string
}
