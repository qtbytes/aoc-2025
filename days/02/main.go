package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part 1
//
//	func handle(x, y int) (ans int) {
//		s := strconv.Itoa(x)
//		var start int
//		if len(s)&1 == 1 {
//			start = int(math.Pow10(len(s) / 2))
//		} else {
//			start, _ = strconv.Atoi(s[:len(s)/2])
//		}
//		for {
//			z, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(start), 2))
//			if z > y {
//				break
//			}
//			if z >= x {
//				ans += z
//			}
//			start++
//		}
//		return ans
//	}

func handle(x, y int) (ans int) {
	s := strconv.Itoa(x)
	seen := make(map[int]bool)
	// repeat times
	for _, t := range []int{2, 3, 5, 7} {
		start, _ := strconv.Atoi(s[:len(s)/t])
		for {
			z, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(start), t))
			if z > y {
				break
			}
			if z >= x {
				// fmt.Println(x, y, z)
				seen[z] = true
			}
			start++
		}
	}
	for x := range seen {
		fmt.Println(x)
		ans += x
	}
	return
}

func main() {
	// f, _ := os.Open("days/02/sample.txt")
	f, _ := os.Open("days/02/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ans := 0
	for scanner.Scan() {
		for s := range strings.SplitSeq(scanner.Text(), ",") {
			s = strings.Trim(s, " ")
			if len(s) == 0 {
				continue
			}
			s := strings.Split(s, "-")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])
			ans += handle(x, y)
		}
	}
	fmt.Println(ans)
}
