package main

import(
	"testing"
)

func TestRepl(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{input: " hello world ",
		expected: []string{"hello", "world"},
		},
		{input: "squirtle SQUIRTLE squirt",
		expected: []string{"squirtle", "squirtle", "squirt"},
		},
		{input: "   bulba  BULBA   bulbasaur",
		expected: []string{"bulba", "bulba", "bulbasaur"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, but got %d", len(c.expected), len(actual))
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word %d to be '%s', but got '%s'", i, expectedWord, word)
				}
			}
		};
	}
}

