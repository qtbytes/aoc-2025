package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
	"strconv"
	"strings"
)

func handle1(lights string, buttons []string) int {
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
		ans += handle1(lights, buttons)
	}
	fmt.Println(ans)
}

func handle2(buttons []string, joltage string) int {
	joltage = joltage[1 : len(joltage)-1]
	var target []int
	for num := range strings.SplitSeq(joltage, ",") {
		num, _ := strconv.Atoi(num)
		target = append(target, num)
	}

	var pairs [][]int
	for _, button := range buttons {
		var pair []int
		for num := range strings.SplitSeq(button[1:len(button)-1], ",") {
			num, _ := strconv.Atoi(num)
			pair = append(pair, num)
		}
		pairs = append(pairs, pair)
	}

	var dfs func(i int, cnt []int) int
	dfs = func(i int, cnt []int) int {
		if i == len(pairs) {
			if slices.Equal(cnt, target) {
				return 0
			} else {
				return 100
			}
		}
		ans := dfs(i+1, cnt)
		for j := 1; ; j++ {
			ok := true
			newCnt := slices.Clone(cnt)
			for _, x := range pairs[i] {
				newCnt[x] += j
				if newCnt[x] > target[x] {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
			ans = min(ans, j+dfs(i+1, newCnt))
		}
		return ans
	}

	cnt := make([]int, len(target))
	ans := dfs(0, cnt)
	fmt.Println(target, pairs, ans)
	return ans
}

func part2(scanner *bufio.Scanner) {
	ans := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		n := len(line)
		buttons := line[1 : n-1]
		cnt := line[n-1]
		ans += handle2(buttons, cnt)
	}
	fmt.Println(ans)
}
func main() {
	// f, _ := os.Open("days/10/sample.txt")
	f, _ := os.Open("days/10/input.txt")
	scanner := bufio.NewScanner(f)
	// part1(scanner)
	part2(scanner)
}
