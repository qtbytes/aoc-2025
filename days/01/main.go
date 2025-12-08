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
		if dir == 'L' {
			pos -= num
		} else {
			pos += num
		}
		pos = (pos + 100) % 100
		if pos == 0 {
			ans++
		}
	}
	fmt.Println(ans)

}
