package lisp

import (
	"fmt"
	"strconv"
)

func handle(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("error: %s", r)
	}
}

func mustParseInt(a Expr) int64 {
	num, err := strconv.ParseInt(a.String(), 10, 64)
	if err != nil {
		panic("could not read integer from given value:" + a.String())
	}

	return num
}
