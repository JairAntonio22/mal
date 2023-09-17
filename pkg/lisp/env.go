package lisp

var ReplEnv Env = Env{
	"+": sum,
	"-": sub,
	"*": mul,
	"/": div,
}

func sum(args ...Expr) Expr {
	ans, _ := args[0].(Number)

	for _, arg := range args[1:] {
		n, _ := arg.(Number)
		ans += n
	}

	return Number(ans)
}

func sub(args ...Expr) Expr {
	ans, _ := args[0].(Number)

	for _, arg := range args[1:] {
		n, _ := arg.(Number)
		ans -= n
	}

	return Number(ans)
}

func mul(args ...Expr) Expr {
	ans, _ := args[0].(Number)

	for _, arg := range args[1:] {
		n, _ := arg.(Number)
		ans *= n
	}

	return Number(ans)
}

func div(args ...Expr) Expr {
	ans, _ := args[0].(Number)

	for _, arg := range args[1:] {
		n, _ := arg.(Number)
		ans /= n
	}

	return Number(ans)
}
