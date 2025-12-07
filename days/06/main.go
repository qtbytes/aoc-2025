package main

import (
	"bufio"
	"fmt"
	"os"
)

func add(nums []int) int {
	ans := 0
	for _, x := range nums {
		ans += x
	}
	return ans
}

func mul(nums []int) int {
	ans := 1
	for _, x := range nums {
		ans *= x
	}
	return ans
}

// Part 1
//
// func main() {
// 	f, _ := os.Open("days/06/input.txt")
//  defer f.Close()
// 	scanner := bufio.NewScanner(f)
// 	n := 1000 // read from the input file
// 	a := make([][]int, n)
// 	ans := 0

// 	for scanner.Scan() {
// 		input := scanner.Text()
// 		t := strings.Fields(input)
// 		if t[0] == "*" || t[0] == "+" {
// 			for i, op := range t {
// 				if op == "+" {
// 					ans += add(&a[i])
// 				} else {
// 					ans += mul(&a[i])
// 				}
// 			}
// 		} else {
// 			for i, s := range t {
// 				x, _ := strconv.Atoi(s)
// 				a[i] = append(a[i], x)
// 			}

// 		}
// 	}
// 	fmt.Println(ans)
// }

func main() {
	f, _ := os.Open("days/06/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var board [][]byte

	maxL := 0
	for scanner.Scan() {
		// notice, here can't use `scanner.Bytes()`
		// it reuses scanner inner []byte, don't allocate
		line := scanner.Text()
		board = append(board, []byte(line))
		maxL = max(maxL, len(line))
	}

	n := len(board)
	ops := board[n-1]
	board = board[:n-1]

	// find least significant digit
	low := make([]int, maxL)
	for i := range low {
		low[i] = -1
	}
	for i, row := range board {
		for j, ch := range row {
			if ch != byte(' ') {
				low[j] = i
			}
		}
	}

	nums := make([]int, len(low))
	for j := range maxL {
		for i := range low[j] + 1 {
			var x int
			if board[i][j] == byte(' ') {
				x = 0
			} else {
				x = int(board[i][j] - byte('0'))
			}
			nums[j] = nums[j]*10 + x
		}
	}

	ans := 0
	for i := 0; i < maxL; i++ {
		r := i + 1
		for r < maxL && low[r] != -1 {
			r++
		}
		cur := nums[i:r]
		// fmt.Println(cur)
		if ops[i] == byte('+') {
			ans += add(cur)
		} else {
			ans += mul(cur)
		}
		i = r
	}
	fmt.Println(ans)
}
