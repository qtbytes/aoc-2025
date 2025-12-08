package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("days/03/input.txt")
	// f, _ := os.Open("days/03/sample.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		n := len(line)
		var nums []int
		for _, x := range line {
			nums = append(nums, int(x-'0'))
		}
		maxBit := nums[n-1]
		res := 0
		for i := n - 2; i >= 0; i-- {
			res = max(res, nums[i]*10+maxBit)
			maxBit = max(maxBit, nums[i])
		}
		ans += res
	}
	fmt.Println(ans)
}
