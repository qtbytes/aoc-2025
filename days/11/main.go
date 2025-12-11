package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// f, _ := os.Open("days/11/sample2.txt")
	f, _ := os.Open("days/11/input.txt")
	scanner := bufio.NewScanner(f)
	// part1(scanner)
	part2(scanner)
}

func part1(scanner *bufio.Scanner) {
	g := make(map[string][]string)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		start := line[0]
		for end := range strings.SplitSeq(line[1], " ") {
			if len(end) == 0 {
				continue
			}
			g[start] = append(g[start], end)
		}
	}

	memo := make(map[string]int)
	var dfs func(start string) int
	dfs = func(start string) int {
		if start == "out" {
			return 1
		}
		if ans, ok := memo[start]; ok {
			return ans
		}
		ans := 0
		for _, end := range g[start] {
			ans += dfs(end)
		}
		memo[start] = ans
		return ans
	}

	fmt.Println(g)
	fmt.Println(dfs("you"))
}

func part2(scanner *bufio.Scanner) {
	g := make(map[string][]string)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		start := line[0]
		for end := range strings.SplitSeq(line[1], " ") {
			if len(end) == 0 {
				continue
			}
			g[start] = append(g[start], end)
		}
	}

	memo := make(map[string][4]int)
	var dfs func(start string) [4]int
	// lowbit 1 means contains "dac"
	// highbit 1 means contains "fft"
	// [00,01,10,11]
	dfs = func(start string) [4]int {
		if start == "out" {
			return [4]int{1}
		}
		if ans, ok := memo[start]; ok {
			return ans
		}
		ans := [4]int{}
		for _, end := range g[start] {
			sub := dfs(end)
			for i, x := range sub {
				ans[i] += x
			}
			if start == "dac" {
				ans[3] += sub[2]
				ans[1] += sub[0]
			}
			if start == "fft" {
				ans[3] += sub[1]
				ans[2] += sub[0]
			}
		}
		memo[start] = ans
		return ans
	}

	fmt.Println(g)
	ans := dfs("svr")
	fmt.Println(ans)
}
