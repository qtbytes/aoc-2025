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
	scanner.Scan()
	first := scanner.Text()
	n := len(first)
	i := strings.Index(first, "S")
	prev := make(map[int]bool)
	prev[i] = true
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		curr := make(map[int]bool)
		for i, ok := range prev {
			if !ok || i < 0 || i >= n {
				continue
			}
			if line[i] == '^' {
				ans++
				curr[i-1] = true
				curr[i+1] = true
			} else {
				curr[i] = true
			}
		}
		prev = curr
	}
	fmt.Println(ans)
}
