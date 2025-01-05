package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " pikachu bulbasaur mewtwo ",
			expected: []string{"pikachu", "bulbasaur", "mewtwo"},
	},
		{
			input: " charmander charizard mew ",
			expected: []string{"charmander", "charizard", "mew"},
	},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input %v, expected %v but got %v", c.input, expectedWord, word)
			}
		}
	}
}