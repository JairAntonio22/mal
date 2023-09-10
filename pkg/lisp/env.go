package lisp

import "strconv"

type Func func(Atom, Atom) Atom

var repl_env map[Atom]Func = map[Atom]Func{
	"+": sum,
	"-": sum,
	"*": sum,
	"/": sum,
}

func sum(a1, a2 Atom) Atom {
	n1, err := strconv.ParseInt(a1.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	n2, err := strconv.ParseInt(a2.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	ans := n1 + n2
	literal := strconv.FormatInt(ans, 10)
	return Atom(literal)
}

func sub(a1, a2 Atom) Atom {
	n1, err := strconv.ParseInt(a1.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	n2, err := strconv.ParseInt(a2.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	ans := n1 - n2
	literal := strconv.FormatInt(ans, 10)
	return Atom(literal)
}

func prod(a1, a2 Atom) Atom {
	n1, err := strconv.ParseInt(a1.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	n2, err := strconv.ParseInt(a2.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	ans := n1 * n2
	literal := strconv.FormatInt(ans, 10)
	return Atom(literal)
}

func div(a1, a2 Atom) Atom {
	n1, err := strconv.ParseInt(a1.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	n2, err := strconv.ParseInt(a2.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	ans := n1 / n2
	literal := strconv.FormatInt(ans, 10)
	return Atom(literal)
}
