package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple": {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		"empty": {
			input:    "",
			expected: []string{""},
		},
		"all spaces": {
			input:    "     ",
			expected: []string{""},
		},
		"single word": {
			input:    "hello",
			expected: []string{"hello"},
		},
		"leading and trailing spaces": {
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		"triple spaces between words": {
			input:    "hello   world",
			expected: []string{"hello", "world"},
		},
		"quadruple spaces between words": {
			input:    "hello    world",
			expected: []string{"hello", "world"},
		},
		"quintuple spaces between words": {
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		"tabs not trimmed": {
			input:    "\t\thello\t\t",
			expected: []string{"\t\thello\t\t"},
		},
		"unicode preserved": {
			input:    "éxito",
			expected: []string{"éxito"},
		},
		"numbers with spaces": {
			input:    " pikachu 25 ",
			expected: []string{"pikachu", "25"},
		},
		"punctuation preserved": {
			input:    "hello, world!",
			expected: []string{"hello,", "world!"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}

}

func ExampleCleanInput() {
	result := CleanInput("  hello  world  ")
	fmt.Println(result)
	// Output: [hello world]
}

func ExampleCleanInput_caseInsensitive() {
	result := CleanInput("  HELLO  WoRlD  ")
	fmt.Println(result)
	// Output: [hello world]
}

func ExampleCleanInput_multipleSpaces() {
	result := CleanInput("hello   world")
	fmt.Println(result)
	// Output: [hello world]
}

func ExampleCleanInput_singleWord() {
	result := CleanInput("  PIKACHU  ")
	fmt.Println(result)
	// Output: [pikachu]
}
