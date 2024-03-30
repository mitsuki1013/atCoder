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
