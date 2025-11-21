package main

import "testing"	

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: 	"  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
        	input: "mIxed CAse",
        	expected: []string{"mixed", "case"},
    	},
    	{
        	input: "whisper SHOUT",
        	expected: []string{"whisper", "shout"},
    	},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths of actual and expected output don't match. Actual: %d, Expected: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words at indes %d don't match: Actual: %s, Expected: %s", i, word, expectedWord)
			}
		}
	}
}