package test

import (
	"fmt"
	"testing"

	"github.com/JairAntonio22/mal/pkg/lisp"
)

type testCase struct {
	given string
	want  string
}

func testRead(t *testing.T, testCases []testCase) {
	for _, tc := range testCases {
		name := fmt.Sprintf("%s", tc.given)

		t.Run(name, func(t *testing.T) {
			expr, err := lisp.Read(tc.given)
			if err != nil {
				t.Fatalf("Could not read '%s'", tc.given)
			}

			if got := expr.String(); got != tc.want {
				t.Errorf("Read(%s) = %s; want %s", tc.given, got, tc.want)
			}
		})
	}
}

func TestReadSymbol(t *testing.T) {
	testRead(t, []testCase{
		{
			given: "+",
			want:  "+",
		},
		{
			given: "abc",
			want:  "abc",
		},
		{
			given: "abc5",
			want:  "abc5",
		},
		{
			given: "abc-def",
			want:  "abc-def",
		},
	})
}

func TestReadList(t *testing.T) {
	testRead(t, []testCase{
		{
			given: "(+ 1 2)",
			want:  "(+ 1 2)",
		},
		{
			given: "()",
			want:  "()",
		},
		{
			given: "( )",
			want:  "()",
		},
		{
			given: "(nil)",
			want:  "(nil)",
		},
		{
			given: "((3 4))",
			want:  "((3 4))",
		},
		{
			given: "(+ 1 (+ 2 3))",
			want:  "(+ 1 (+ 2 3))",
		},
		{
			given: "  ( +   1   (+   2 3   )   )  ",
			want:  "(+ 1 (+ 2 3))",
		},
		{
			given: "(* 1 2)",
			want:  "(* 1 2)",
		},
		{
			given: "(** 1 2)",
			want:  "(** 1 2)",
		},
		{
			given: "(* -3 6)",
			want:  "(* -3 6)",
		},
		{
			given: "(()())",
			want:  "(() ())",
		},
	})
}

func TestIgnoreCommas(t *testing.T) {
	testRead(t, []testCase{
		{
			given: "(1 2, 3,,,,),,",
			want:  "(1 2 3)",
		},
	})
}

func TestConsts(t *testing.T) {
	testRead(t, []testCase{
		{
			given: "nil",
			want:  "nil",
		},
		{
			given: "true",
			want:  "true",
		},
		{
			given: "false",
			want:  "false",
		},
	})
}

func TestReadStrings(t *testing.T) {
	testRead(t, []testCase{
		{
			given: `"abc"`,
			want:  `"abc"`,
		},
		{
			given: `   "abc"   "`,
			want:  `"abc"`,
		},
		{
			given: `"abc (with parens)"`,
			want:  `"abc (with parens)"`,
		},
		{
			given: `"abc\"def"`,
			want:  `"abc\"def"`,
		},
		{
			given: `""`,
			want:  `""`,
		},
		{
			given: `"\\"`,
			want:  `"\\"`,
		},
		{
			given: `"\\\\\\\\\\\\\\\\\\"`,
			want:  `"\\\\\\\\\\\\\\\\\\"`,
		},
	})
}
