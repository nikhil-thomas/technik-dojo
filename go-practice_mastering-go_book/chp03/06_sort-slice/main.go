package main

import (
	"fmt"
	"sort"
)

type person struct {
	name   string
	height int
	weight int
}

func main() {
	ps := make([]person, 0)
	ps = append(ps, person{"B", 180, 80})
	ps = append(ps, person{"A", 170, 70})
	ps = append(ps, person{"D", 200, 100})
	ps = append(ps, person{"C", 190, 90})

	fmt.Println(ps)

	fmt.Printf("\n height ascending")
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].height < ps[j].height
	})
	fmt.Println(ps)

	fmt.Printf("\n height descending")
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].height > ps[j].height
	})
	fmt.Println(ps)
}
