package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, _ := os.Open("days/04/sample.txt")
	f, _ := os.Open("days/04/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var board []string
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}
	fmt.Println(board)
	ans := 0
	m, n := len(board), len(board[0])
	for i, s := range board {
		for j, ch := range s {
			if ch != '@' {
				continue
			}
			cnt := -1 // will count self again
			for x := max(0, i-1); x < min(m, i+2); x++ {
				for y := max(0, j-1); y < min(n, j+2); y++ {
					if board[x][y] == '@' {
						cnt++
					}
				}
			}
			if cnt < 4 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
