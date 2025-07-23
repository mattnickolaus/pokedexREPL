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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO",
			expected: []string{"hello"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual_result := cleanInput(c.input)

		if len(actual_result) != len(c.expected) {
			t.Errorf("%v != %v\nLengths: %v != %v", actual_result, c.expected, len(actual_result), len(c.expected))
		}

		for i, _ := range actual_result {
			word := actual_result[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%v != %v\n", word, expectedWord)
			}
		}
	}
}
