package lisp

import (
	"fmt"
	"strconv"
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

type Number int

func (n Number) isExpr() {}

func (n Number) String() string {
	return strconv.Itoa(int(n))
}

type List []Expr

func (l List) isExpr() {}

func (l List) String() string {
	var sb strings.Builder
	sb.WriteRune('(')

	for i, e := range l {
		sb.WriteString(e.String())

		if i < len(l)-1 {
			sb.WriteRune(' ')
		}
	}

	sb.WriteRune(')')
	return sb.String()
}

type Func func(...Expr) Expr

func (f Func) isExpr() {}

func (f Func) String() string { return "" }

type Env map[Atom]Func
