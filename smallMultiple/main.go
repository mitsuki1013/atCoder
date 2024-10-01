package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
)

// テスト用に標準入力と標準出力を変数として定義
var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

func main() {
	var K int
	fmt.Fscan(stdin, &K)

	queue := list.New()
	visited := make(map[int]bool)

	queue.PushBack(1)
	visited[1] = true

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		current := front.Value.(int)

		if current%K == 0 {
			sumOfDigits := 0
			s := strconv.Itoa(current)
			for _, c := range s {
				sumOfDigits += int(c - '0')
			}
			fmt.Fprint(stdout, sumOfDigits, "\n")
			return
		}

		for i := 0; i < 10; i++ {
			next := current*10 + i
			nextModK := next % K
			if !visited[nextModK] {
				visited[nextModK] = true
				queue.PushBack(next)
			}
		}
	}
}
