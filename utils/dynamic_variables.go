package utils

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func CalcExpression(expression string) (string, error) {
	handler, _ := govaluate.NewEvaluableExpression(expression)
	result, err := handler.Evaluate(nil)

	return fmt.Sprintf("%v", result), err
}
