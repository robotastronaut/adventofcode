package main

import (
	"testing"
)

func TestReplaceNumberWords(t *testing.T) {
	cases := [][]string{
		{"two1nine", "219"},
		{"eightwothree", "823"},
		{"abcone2threexyz", "abc123xyz"},
		{"xtwone3four", "x2134"},
		{"4nineeightseven2", "49872"},
		{"zoneight234", "z18234"},
		{"7pqrstsixteen", "7pqrst6teen"},
		{"195one", "1951"}}

	for _, c := range cases {
		res := replaceNumberWords(c[0])
		if res != c[1] {
			t.Errorf("Error, got: %s, want: %s", res, c[1])
		}
	}
}
