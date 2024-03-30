package main

import "testing"

func Test(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "ケース1",
			input: [][]int{
				{1, 5},
				{2, 4},
				{3, 6},
			},
			expected: 3,
		},
		{
			name: "ケース2",
			input: [][]int{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			},
			expected: 27,
		},
		{
			name: "ケース3",
			input: [][]int{
				{3, 14, 159, 2, 6, 53},
				{58, 9, 79, 323, 84, 6},
				{2643, 383, 2, 79, 50, 288},
			},
			expected: 87,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := CountPatterns(c.input)
			if CountPatterns(c.input) != c.expected {
				t.Errorf("期待するoutputではない\ngot: %d\nwant: %d", result, c.expected)
			}
		})
	}
}
