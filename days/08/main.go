package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Pos struct {
	x, y, z int
}
type Pair struct {
	dist, i, j int
}

func calcDistance(p1, p2 Pos) int {
	dx, dy, dz := p1.x-p2.x, p1.y-p2.y, p1.z-p2.z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	// f, _ := os.Open("days/08/sample.txt")
	// k := 10
	f, _ := os.Open("days/08/input.txt")
	k := 1000
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var q []Pos
	for scanner.Scan() {
		var a []int
		for x := range strings.SplitSeq(scanner.Text(), ",") {
			x, _ := strconv.Atoi(x)
			a = append(a, x)
		}
		q = append(q, Pos{a[0], a[1], a[2]})
	}
	n := len(q)
	// fmt.Println(q, len(q))

	var ds []Pair
	for i := range len(q) {
		for j := range i {
			ds = append(ds, Pair{calcDistance(q[i], q[j]), i, j})
		}
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].dist < ds[j].dist
	})
	// fmt.Println(ds)

	// Union Find
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}
	union := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		size[fx] += size[fy]
		fa[fy] = fx
	}

	for i := range k {
		d := ds[i]
		union(d.i, d.j)
	}

	group := make([]int, n)
	for x := range n {
		f := find(x)
		group[f] = size[f]
	}
	sort.Ints(group)
	slices.Reverse(group)
	fmt.Println(group[0] * group[1] * group[2])
}
