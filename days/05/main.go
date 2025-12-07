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

type Pair struct {
	l, r int
}

func main() {
	f, _ := os.Open("days/05/input.txt")
	scanner := bufio.NewScanner(f)
	var a []Pair
	ans := 0
	for scanner.Scan() {
		input := scanner.Text()
		if strings.Contains(input, "-") {
			t := strings.Split(input, "-")
			start, _ := strconv.Atoi(t[0])
			end, _ := strconv.Atoi(t[1])

			i := sort.Search(len(a), func(i int) bool {
				return a[i].r >= start
			})

			if i == len(a) {
				a = append(a, Pair{start, end})
			} else if a[i].l > end {
				a = slices.Insert(a, i, Pair{start, end})
			} else {
				j := i
				for ; j < len(a); j++ {
					if end >= a[j].l {
						start = min(start, a[j].l)
						end = max(end, a[j].r)
					} else {
						break
					}
				}
				a[i] = Pair{start, end}
				a = slices.Delete(a, i+1, j)
			}
		} else {
			x, _ := strconv.Atoi(input)
			i := sort.Search(len(a), func(i int) bool { return a[i].r > x })
			if i < len(a) && a[i].l <= x && x <= a[i].r {
				ans++
			}

		}
	}
	fmt.Println(ans)
	res := 0
	for _, p := range a {
		res += p.r - p.l + 1
	}
	fmt.Println(res)
}
