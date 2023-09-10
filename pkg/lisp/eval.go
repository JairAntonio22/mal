package lisp

import (
	"fmt"
)

func Eval(input Expr) (expr Expr, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error ocurred while evaluating expr: %s", r)
		}
	}()

	switch v := input.(type) {
	case Atom:
		return v, nil

	case List:
		return v, nil

	default:
		panic("Unexpected type given to evaluate")
	}
}
