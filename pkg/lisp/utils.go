package lisp

import "fmt"

func handle(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("error: %s", r)
	}
}
