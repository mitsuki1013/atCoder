package main

import (
	"bytes"
	"testing"
)

// ユニットテスト用の関数
func TestMainFunction(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"6", "3\n"},
		{"41", "5\n"},
		{"79992", "36\n"},
	}

	for _, testCase := range testCases {
		t.Run("K="+testCase.input, func(t *testing.T) {
			// 入力を設定
			input := bytes.NewBufferString(testCase.input)

			// 出力をキャプチャ
			var output bytes.Buffer

			// 標準入力と出力を設定
			stdin = input
			stdout = &output

			// main関数を呼び出し
			main()

			// 出力を検証
			if output.String() != testCase.expected {
				t.Errorf("expected %s but got %s", testCase.expected, output.String())
			}
		})
	}
}
