package lisp

func Eval(input Expr, env Env) (expr Expr, err error) {
	defer handle(&err)

	list, ok := input.(List)

	if !ok {
		return evalAst(input, env), nil
	}

	if len(list) == 0 {
		return input, nil
	}

	list = evalAst(list, env).(List)

	f := list[0].(Func)
	args := list[1:]

	return f(args...), nil
}

func evalAst(input Expr, env Env) Expr {
	switch input := input.(type) {
	case Atom:
		val, exists := env[input]
		if !exists {
			panic("undefined symbol" + input)
		}

		return val

	case List:
		list := List{}

		for _, subexpr := range input {
			result, _ := Eval(subexpr, env)

			list = append(list, result)
		}

		return list

	default:
		return input
	}
}
