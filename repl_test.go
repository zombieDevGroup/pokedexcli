package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string // adding a name field helps identify which test case failed
		input    string
		expected []string
	}{
		{
			name:     "basic case with spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed case with pokemon",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name:     "single word",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("got len %d, want len %d for input '%s'",
				len(actual), len(c.expected), c.input)
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// Check each word
			if word != expectedWord {
				t.Errorf("got '%s', want '%s' for input '%s'",
					word, expectedWord, c.input)
				return
			}
		}

	}

}
