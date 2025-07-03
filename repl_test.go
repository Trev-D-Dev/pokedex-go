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
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %v  actual length: %v", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expWord := c.expected[i]
			if word != expWord {
				t.Errorf("expected word: %v  actual word: %v", expWord, word)
			}
		}
	}
}
