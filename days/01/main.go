package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// f, _ := os.Open("days/01/sample.txt")
	f, _ := os.Open("days/01/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ans := 0
	pos := 50
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		num, _ := strconv.Atoi(line[1:])
		// for each 100 move, must reach 0 one times
		if num >= 100 {
			ans += num / 100
			num %= 100
		}
		ori := pos
		if dir == 'L' {
			pos -= num
		} else {
			pos += num
		}
		if 0 < pos && pos < 100 {
			continue
		}
		if ori != 0 {
			ans++
		}
		pos = (pos%100 + 100) % 100
	}
	fmt.Println(ans)
}
