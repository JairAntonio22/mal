package lisp

import (
	"fmt"
	"strings"
)

type Expr interface {
	isExpr()
	fmt.Stringer
}

type Atom string

func (a Atom) isExpr() {}

func (a Atom) String() string {
	return string(a)
}

type List []Expr

func (l List) isExpr() {}

func (l List) String() string {
	var sb strings.Builder
	sb.WriteRune('(')

	for i, expr := range l {
		sb.WriteString(expr.String())

		if i < len(l)-1 {
			sb.WriteRune(' ')
		}
	}

	sb.WriteRune(')')
	return sb.String()
}
