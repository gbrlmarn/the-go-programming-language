// Exericse 7.3: Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4) that reveals the sequence of values in the tree.
package main

import (
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add appends input value to tree and returns tree
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// String returns the set as a string of the form "{1 2 3}".
func (t *tree) String() string {
	var s string
	if t == nil {
		return s
	}
	s += t.left.String()
	s += strconv.Itoa(t.value)
	s += t.right.String()
	return s
}
