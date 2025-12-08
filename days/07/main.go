package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// f, _ := os.Open("days/07/sample.txt")
	f, _ := os.Open("days/07/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var board []string
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}
	// fmt.Println(board)
	m, n := len(board), len(board[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if board[i][j] == '.' {
			*p = dfs(i+1, j)
		} else {
			*p = dfs(i+1, j-1) + dfs(i+1, j+1)
		}
		return *p
	}

	j := strings.Index(board[0], "S")
	fmt.Println(dfs(1, j))

	// Part 1
	//
	// scanner.Scan()
	// first := scanner.Text()
	// n := len(first)
	// i := strings.Index(first, "S")
	// prev := make(map[int]bool)
	// prev[i] = true
	// ans := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	curr := make(map[int]bool)
	// 	for i, ok := range prev {
	// 		if !ok || i < 0 || i >= n {
	// 			continue
	// 		}
	// 		if line[i] == '^' {
	// 			ans++
	// 			curr[i-1] = true
	// 			curr[i+1] = true
	// 		} else {
	// 			curr[i] = true
	// 		}
	// 	}
	// 	prev = curr
	// }
	// fmt.Println(ans)
}
