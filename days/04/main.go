package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	i, j int
}

type Item struct {
	i, j, cnt int
}

func main() {
	// f, _ := os.Open("days/04/sample.txt")
	f, _ := os.Open("days/04/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var board [][]rune
	for scanner.Scan() {
		board = append(board, []rune(scanner.Text()))
	}
	m, n := len(board), len(board[0])
	var q []Item
	adj := make(map[Pair]int) // count
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
				q = append(q, Item{i, j, cnt})
			}
			adj[Pair{i, j}] = cnt
		}
	}
	ans := 0
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		i, j := item.i, item.j
		if board[i][j] != '@' {
			continue
		}
		ans++
		board[i][j] = 'x'
		for x := max(0, i-1); x < min(m, i+2); x++ {
			for y := max(0, j-1); y < min(n, j+2); y++ {
				if x == i && y == j {
					continue
				}
				if board[x][y] == '@' {
					adj[Pair{x, y}]--
					if adj[Pair{x, y}] < 4 {
						q = append(q, Item{x, y, adj[Pair{x, y}]})
					}
				}
			}
		}
	}
	fmt.Println(ans)

	// Part 1
	//
	// ans := 0
	// m, n := len(board), len(board[0])
	// for i, s := range board {
	// 	for j, ch := range s {
	// 		if ch != '@' {
	// 			continue
	// 		}
	// 		cnt := -1 // will count self again
	// 		for x := max(0, i-1); x < min(m, i+2); x++ {
	// 			for y := max(0, j-1); y < min(n, j+2); y++ {
	// 				if board[x][y] == '@' {
	// 					cnt++
	// 				}
	// 			}
	// 		}
	// 		if cnt < 4 {
	// 			ans++
	// 		}
	// 	}
	// }
	// fmt.Println(ans)
}
