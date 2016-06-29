package namedregexp

import (
	"fmt"
	"testing"
)

type input struct {
	pattern, haystack string
}

type expected struct {
	result map[string]string
	err    error
}

type tcase struct {
	input    input
	expected expected
}

func TestNamedGroups(t *testing.T) {
	args := []tcase{
		tcase{
			input: input{
				pattern:  `(?P<first>\d+)\.(\d+).(?P<second>\d+)`,
				haystack: "1234.5678.9",
			},
			expected: expected{map[string]string{"first": "1234", "second": "9"}, nil},
		},
	}

	for _, a := range args {
		out, err := FindNamedStringSubmatch(a.input.pattern, a.input.haystack)
		if err != a.expected.err {
			t.Errorf("Error is %s,expected %s", err, a.expected.err)
		}

		if !mapEqual(out, a.expected.result) {
			t.Errorf("Got result %+v expected %+v", out, a.expected.result)
		}
	}
}

func mapEqual(x, y map[string]string) bool {
	for k, v := range x {
		if val, ok := y[k]; ok {
			fmt.Printf("k: %s, v: %s, otherv: %s\n", k, v, val)
			if val != v {

				return false
			}
		} else {
			return false
		}
	}
	return true
}

// func mapHasKey(subject map[string]string, key string) bool {
// 	for k, v := range subject {
// 		if k == key {
// 			return true
// 		}
// 	}
// 	return false
// }
