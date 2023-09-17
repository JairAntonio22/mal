package test

import (
	"fmt"
	"testing"

	"github.com/JairAntonio22/mal/pkg/lisp"
)

func testEval(t *testing.T, testCases []testCase) {
	for _, tc := range testCases {
		name := fmt.Sprintf("%s", tc.given)

		t.Run(name, func(t *testing.T) {
			input, _ := lisp.Read(tc.given)

			expr, err := lisp.Eval(input, lisp.ReplEnv)
			if err != nil {
				t.Fatalf("Could not read '%s'", tc.given)
			}

			if got := expr.String(); got != tc.want {
				t.Errorf("Eval(%s) = %s; want %s", tc.given, got, tc.want)
			}
		})
	}
}

func TestEvalArithmetic(t *testing.T) {
	testEval(t, []testCase{
		{
			given: "(+ 1 2)",
			want:  "3",
		},
		{
			given: "(+ 5 (* 2 3))",
			want:  "11",
		},
		{
			given: "(- (+ 5 (* 2 3)) 3)",
			want:  "8",
		},
		{
			given: "(/ (- (+ 5 (* 2 3)) 3) 4)",
			want:  "2",
		},
		{
			given: "(/ (- (+ 515 (* 87 311)) 302) 27)",
			want:  "1010",
		},
		{
			given: "(* -3 6)",
			want:  "-18",
		},
		{
			given: "(/ (- (+ 515 (* -87 311)) 296) 27)",
			want:  "-994",
		},
		{
			given: "()",
			want:  "()",
		},
	})
}
