package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func segmentIntersectRectangle(a, b, p1, p2 Pos) bool {
	minX, maxX := p1.x, p2.x
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := p1.y, p2.y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	x1, y1 := a.x, a.y
	x2, y2 := b.x, b.y

	if max(x1, x2) < minX || min(x1, x2) > maxX ||
		max(y1, y2) < minY || min(y1, y2) > maxY {
		return false
	}

	if x1 <= minX && x2 <= minX {
		return false
	}
	if x1 >= maxX && x2 >= maxX {
		return false
	}
	if y1 <= minY && y2 <= minY {
		return false
	}
	if y1 >= maxY && y2 >= maxY {
		return false
	}

	return true
}

type Pair struct {
	area, i, j int
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

// can't handle situation like sample2
func part2(scanner *bufio.Scanner) {
	var a []Pos
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[1])
		y, _ := strconv.Atoi(line[0])
		a = append(a, Pos{x, y})
	}

	var areas []Pair

	for i := range a {
		for j := range i {
			area := cacl(a[i], a[j])
			areas = append(areas, Pair{area, i, j})
		}
	}

	sort.Slice(areas, func(i, j int) bool {
		return areas[i].area > areas[j].area
	})

	check := func(p Pair) bool {
		for i := range a {
			if segmentIntersectRectangle(a[i], a[(i+1)%len(a)], a[p.i], a[p.j]) {
				return false
			}
		}
		return true
	}

	for _, p := range areas {
		if check(p) {
			fmt.Println(p.area)
			return
		}
	}
}

func main() {
	// f, _ := os.Open("days/09/sample2.txt")
	f, _ := os.Open("days/09/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	part2(scanner)

	// var a []Pos

	// ans := 0
	// for scanner.Scan() {
	// 	line := strings.Split(scanner.Text(), ",")
	// 	x, _ := strconv.Atoi(line[1])
	// 	y, _ := strconv.Atoi(line[0])
	// 	p := Pos{x, y}
	// 	for _, p2 := range a {
	// 		ans = max(ans, cacl(p, p2))
	// 	}
	// 	a = append(a, p)
	// }
	// fmt.Println(ans)
	// sort.Slice(a, func(i, j int) bool {
	// 	return a[i].x < a[j].x || a[i].x == a[j].x && a[i].y < a[j].y
	// })

}
