package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " test number one",
			expected: []string{"test", "number", "one"},
		},
		{
			input:    "test number two ",
			expected: []string{"test", "number", "two"},
		},
		{
			input:    "test  number  three",
			expected: []string{"test", "number", "three"},
		},
		{
			input:    "test number 4",
			expected: []string{"test", "number", "4"},
		},
		{
			input:    "Test number 5",
			expected: []string{"test", "number", "5"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Expected Length: %v\nActual Length: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected Word: %s\nActual Word: %s", expectedWord, word)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}

}
