package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   HeLlO  wOrLd   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   A GOOD TESTCASE   ",
			expected: []string{"a", "good", "testcase"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput function not working properly for input: %s, different length", c.input)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput function not working properly for input: %s, different expected word", c.input)
			}
		}
	}
}
