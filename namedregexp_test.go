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

// func TestNamedGroups(t *testing.T) {
// 	args := []tcase{
// 		tcase{
// 			input: input{
// 				pattern:  `(?P<first>\d+)\.(\d+).(?P<second>\d+)`,
// 				haystack: "1234.5678.9",
// 			},
// 			expected: expected{map[string]string{"first": "1234", "second": "9"}, nil},
// 		},
// 	}
//
// 	for _, a := range args {
// 		out, err := FindNamedStringSubmatch(a.input.pattern, a.input.haystack)
// 		if err != a.expected.err {
// 			t.Errorf("Error is %s,expected %s", err, a.expected.err)
// 		}
//
// 		if !mapEqual(out, a.expected.result) {
// 			t.Errorf("Got result %+v expected %+v", out, a.expected.result)
// 		}
// 	}
// }

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

func TestAllNamedGroups(t *testing.T) {
	allnamedtests := []struct {
		sample      string
		pattern     string
		errexpected bool
		names       []string
	}{
		{
			sample:      "Mon - Fri: 1000-1900, Sat: 1000-1700, Sun: 1100-1700",
			pattern:     `(?P<days>[\w\s-]+).*?:.*?(?P<openh>[\d]{2})(?P<openm>[\d]{2})-(?P<closeh>[\d]{2})(?P<closem>[\d]{2})`,
			errexpected: false,
			names:       []string{"days", "openh", "openm", "closeh", "closem"},
		},
		{
			sample:      "",
			pattern:     `(?p<miao>[\s]+)`,
			errexpected: true,
			names:       []string{},
		},
		{
			sample:      "ads23",
			pattern:     `(?P<chars>[\s]+)(?P<nums>[\d]+)`,
			errexpected: false,
			names:       []string{"chars", "nums"},
		},
	}

	for _, tt := range allnamedtests {
		match, err := FindAllNamedStringSubmatch(tt.pattern, tt.sample, -1)

		goterr := err != nil
		if goterr != tt.errexpected {
			if err != nil {
				t.Errorf("got err %s, expected %v", err, tt.errexpected)
			} else {
				t.Errorf("got no err, expected %v", tt.errexpected)
			}
		}

		for _, m := range match {
			if len(m) != len(tt.names) {
				t.Errorf("matched names cound missmatch given=%d wanted=%d", len(m), len(tt.names))
			}

			for _, n := range tt.names {
				if _, ok := m[n]; !ok {
					t.Errorf("expected to find %s but did not", n)
				}
			}
		}
	}

}
