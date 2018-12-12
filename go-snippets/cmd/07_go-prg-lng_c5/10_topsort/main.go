package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topsort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topsort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var vistAll func(items []string)
	vistAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				vistAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	vistAll(keys)
	return order
}
