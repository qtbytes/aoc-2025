package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func cacl(p1, p2 Pos) int {
	return (1 + abs(p1.x-p2.x)) * (1 + abs(p1.y-p2.y))
}

func main() {
	// f, _ := os.Open("days/09/example.txt")
	f, _ := os.Open("days/09/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var a []Pos

	ans := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[1])
		y, _ := strconv.Atoi(line[0])
		p := Pos{x, y}
		for _, p2 := range a {
			ans = max(ans, cacl(p, p2))
		}
		a = append(a, p)
	}
	fmt.Println(ans)
	// sort.Slice(a, func(i, j int) bool {
	// 	return a[i].x < a[j].x || a[i].x == a[j].x && a[i].y < a[j].y
	// })

}
