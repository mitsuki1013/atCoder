package main

import (
	"container/list"
	"fmt"
	"github.com/thoas/go-funk"
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

		isMultipleOfK := current%K == 0

		if isMultipleOfK {
			fmt.Fprint(stdout, calculateSumOfDigits(current), "\n")
			return
		}

		enqueueNextStates(current, K, visited, queue)
	}
}

func calculateSumOfDigits(number int) int {
	digits := []rune(strconv.Itoa(number))
	return int(funk.Sum(funk.Map(digits, func(r rune) int {
		return int(r - '0')
	})))
}

func enqueueNextStates(current int, K int, visited map[int]bool, queue *list.List) {
	for i := 0; i < 10; i++ {
		next := current*10 + i
		nextModK := next % K
		if !visited[nextModK] {
			visited[nextModK] = true
			queue.PushBack(next)
		}
	}
}
