package main

import (
	"bufio"
	"github.com/thoas/go-funk"
	"log"
	"os"
	"strconv"
	"strings"
)

const numberOfSteps = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfColumn, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("The first line of input must be an int type.\nerr: %v", err)
	}

	altar := make([][]int, 0, numberOfSteps)
	for i := 0; i < numberOfColumn+1; i++ {
		if i == 0 {
			continue
		}

		scanner.Scan()
		row := funk.Map(strings.Split(scanner.Text(), " "), func(s string) int {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("All inputs must be of type int.\nerr: %v", err)
			}
			return i
		}).([]int)

		if len(row) != numberOfColumn {
			log.Fatal("Not the number of columns specified.")
		}

		altar = append(altar, row)
	}
}

func CountPatterns(altar [][]int) int {
	patterns := funk.Reduce(altar, func(acc [][]int, cur []int) [][]int {
		if len(acc) == 0 {
			acc = append(acc, cur)
			return acc
		}

		items := funk.FlatMap(acc[len(acc)-1], func(fixItem int) []int {
			return funk.Filter(cur, func(item int) bool {
				return fixItem < item
			}).([]int)
		}).([]int)

		acc = append(acc, items)
		return acc
	}, [][]int{}).([][]int)

	return len(patterns[len(patterns)-1])
}
