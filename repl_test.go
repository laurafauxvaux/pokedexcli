package main

import (
	"testing"
)

func TestCleanUp(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	} {
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("got %d words, want%d", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("got %q, want %q", word, expectedWord)
			}
		}
	}
}

