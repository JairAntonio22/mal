package lisp

import (
	"fmt"
)

func Eval(input Expr) (expr Expr, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error ocurred while reading expr: %s", r)
		}
	}()

	return input, nil
}
