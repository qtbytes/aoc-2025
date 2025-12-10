package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func handle(lights string, buttons []string) int {
	// we know each button most be pressed once
	target := 0
	lights = lights[1 : len(lights)-1]
	for i, button := range buttons {
		buttons[i] = button[1 : len(button)-1]
	}
	for i, ch := range lights {
		if ch == '#' {
			target |= 1 << i
		}
	}

	n := len(buttons)
	ans := n
	for mask := range 1 << n {
		value := 0
		cur := bits.OnesCount(uint(mask))
		if cur >= ans {
			continue
		}
		for i := range n {
			if mask>>i&1 == 1 {
				for button := range strings.SplitSeq(buttons[i], ",") {
					button, _ := strconv.Atoi(button)
					value ^= (1 << button)
				}
			}
		}
		if value == target {
			ans = min(ans, cur)
		}
	}
	// fmt.Println(lights, target, buttons, ans)
	return ans
}

func part1(scanner *bufio.Scanner) {
	ans := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		n := len(line)
		lights := line[0]
		buttons := line[1 : n-1]
		ans += handle(lights, buttons)
	}
	fmt.Println(ans)
}

func main() {
	// f, _ := os.Open("days/10/sample.txt")
	f, _ := os.Open("days/10/input.txt")
	scanner := bufio.NewScanner(f)
	part1(scanner)
}
