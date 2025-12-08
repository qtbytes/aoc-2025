package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, _ := os.Open("days/03/sample.txt")
	f, _ := os.Open("days/03/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ans := 0
	const N = 12

	for scanner.Scan() {
		line := scanner.Text()
		f := make([]int, N+1)
		for i, x := range line {
			x := int(x - '0')
			for j := min(i+1, len(f)-1); j > 0; j-- {
				f[j] = max(f[j], f[j-1]*10+x)
			}
		}
		ans += f[12]
	}
	fmt.Println(ans)

	// Part 1
	// scanner := bufio.NewScanner(f)
	// ans := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	n := len(line)
	// 	var nums []int
	// 	for _, x := range line {
	// 		nums = append(nums, int(x-'0'))
	// 	}
	// 	maxBit := nums[n-1]
	// 	res := 0
	// 	for i := n - 2; i >= 0; i-- {
	// 		res = max(res, nums[i]*10+maxBit)
	// 		maxBit = max(maxBit, nums[i])
	// 	}
	// 	ans += res
	// }
	// fmt.Println(ans)
}
