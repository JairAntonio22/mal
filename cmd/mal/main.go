package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/JairAntonio22/mal/pkg/lisp"
)

func main() {
	// TODO(Jair): implement cli functionalities
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		expr, err := lisp.Read(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		expr, err = lisp.Eval(expr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(";=>", expr)
	}
}
