package lisp

import (
	"errors"
	"regexp"
	"strconv"
)

func Read(input string) (expr Expr, err error) {
	defer handle(&err)
	tokens := tokenize(input)
	reader := &reader{tokens: tokens}
	expr = readForm(reader)
	return
}

const (
	SpliceUnquote = `~@`
	Symbols       = `[\[\]{}()'` + "`" + `~^@]`
	Strings       = `"(?:\\.|[^\\"])*"?`
	Comments      = `;.*`
	Atoms         = `[^\s\[\]{}('"` + "`" + `,;)]+`
	MatchCriteria = "(" + SpliceUnquote +
		"|" + Symbols +
		"|" + Strings +
		"|" + Comments +
		"|" + Atoms + ")"
)

var re *regexp.Regexp = regexp.MustCompile(MatchCriteria)

func tokenize(input string) []string {
	return re.FindAllString(input, -1)
}

func readForm(reader *reader) Expr {
	token, err := reader.peek()
	if err != nil {
		panic("missing closing parenthesis")
	}

	switch token {
	case "(":
		reader.next()
		return readList(reader)

	case ")":
		panic("unexpected closing parenthesis")

	default:
		reader.next()

		if num, err := strconv.Atoi(token); err == nil {
			return Number(num)
		}

		return Atom(token)
	}
}

func readList(reader *reader) Expr {
	list := List{}

	for {
		token, err := reader.peek()
		if err != nil {
			panic("missing closing parenthesis")
		}

		if token == ")" {
			reader.next()
			break
		}

		expr := readForm(reader)
		list = append(list, expr)
	}

	return list
}

type reader struct {
	tokens []string
	pos    int
}

func (r *reader) next() (string, error) {
	if r.pos < 0 || r.pos >= len(r.tokens) {
		return "", errors.New("EOF reached unexpectedly")
	}

	t := r.tokens[r.pos]
	r.pos++
	return t, nil
}

func (r reader) peek() (string, error) {
	if r.pos < 0 || r.pos >= len(r.tokens) {
		return "", errors.New("EOF reached unexpectedly")
	}

	return r.tokens[r.pos], nil
}
