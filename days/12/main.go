package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// f, _ := os.Open("days/12/sample.txt")
	f, _ := os.Open("days/12/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	part1(scanner)

}

// ref: https://www.youtube.com/watch?v=am-X5j1DVkA
func part1(scanner *bufio.Scanner) {
	var size []int
	var ans int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, ":") && !strings.Contains(line, "x") {
			cnt := 0
			for range 3 {
				scanner.Scan()
				for _, ch := range scanner.Text() {
					if ch == '#' {
						cnt++
					}
				}
			}
			size = append(size, cnt)
		} else {
			line := strings.Split(line, ":")
			rc := strings.Split(line[0], "x")
			row, _ := strconv.Atoi(rc[0])
			col, _ := strconv.Atoi(rc[1])
			total := row * col
			cnt := 0
			for i, x := range strings.Split(strings.TrimSpace(line[1]), " ") {
				x, _ := strconv.Atoi(x)
				cnt += size[i] * x
			}
			if float64(cnt)*1.3 < float64(total) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
