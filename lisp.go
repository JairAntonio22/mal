package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// TODO(Jair): implement cli functionalities
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("user> ")
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		Rep(input)
	}
}

func Rep(input string) {
	Print(Eval(Read(input)))
}

func Read(input string) string {
	return input
}

func Eval(expr string) string {
	return expr
}

func Print(expr string) {
	fmt.Print(expr)
}
