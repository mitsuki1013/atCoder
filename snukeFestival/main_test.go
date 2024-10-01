package main

import "testing"

func TestCountPatterns(t *testing.T) {
	testCases := []struct {
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
		{
			name:     "ケース4 - 空の入力",
			input:    [][]int{},
			expected: 0,
		},
		{
			name: "ケース5 - 単一の要素",
			input: [][]int{
				{42},
			},
			expected: 1,
		},
		{
			name: "ケース7 - 大きな数字",
			input: [][]int{
				{1000000, 2000000},
				{3000000, 4000000},
				{5000000, 6000000},
			},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runTestCase(t, tc.input, tc.expected)
		})
	}
}

func runTestCase(t *testing.T, input [][]int, expected int) {
	result := CountPatterns(input)
	if result != expected {
		t.Errorf("期待するoutputではない\ngot: %d\nwant: %d", result, expected)
	}
}
