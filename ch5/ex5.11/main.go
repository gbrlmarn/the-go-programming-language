// Exercise 5.11: The instructor of linear algebra course decides that calculus is now a prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer orgranization",
	},

	// introduce a loop
	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer orgranization"},
}

func main() {
	err, sorted := topoSort(prereqs)
	if err != nil {
		log.Fatalf("ex5.11: %v", err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) (error, []string) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
			} else {
				cycle := true
				for _, oitem := range order {
					if oitem == item {
						cycle = false
					}
				}
				if cycle {
					return fmt.Errorf("cycle: %s", item)
				}
			}
		}
		return nil
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return err, nil
	}
	return nil, order
}
