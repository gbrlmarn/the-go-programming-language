// Exercise 6.2: Define a variadic (*IntSet).AddAll(...int) method that allows a list of values to be added, such as s.AddAll(1, 2, 3).
// Exercise 6.1: Implement these additional methods:
// func (*IntSet) Len() int // return the number of elements
// func (*IntSet) Remove(x int) // remove x from the set
// func (*IntSet) Clear() // remove all elements from the set
// func (*IntSet) Copy() *IntSet // return a copy of the set
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Return the number of elements
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ { 
			if word&(1<<uint(j)) != 0 {
				len += 1
			}
		}
	}
	return len
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word := x/64
	if(s.Has(x)) {
		s.words = append(s.words[:word], s.words[word+1:]...)
	}
}

// Remove all elements from the set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := &IntSet{}
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	return c
}

//

// AddAll adds input values to the set.
func (s *IntSet) AddAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}


func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(x.Len(), y.Len()) // "3 2"

	fmt.Println(x.String()) // "{1 9 42 144}"
	x.Remove(144)
	fmt.Println(x.String()) // "{1 9 42}"

	fmt.Println(x.String()) // "{1 9 42}"
	x.Clear()
	fmt.Println(x.String()) // "{}"

	z := y.Copy()
	fmt.Println(z.String()) // "{9 42}"

	z.AddAll(43, 44)
	fmt.Println(z.String()) // "{9 42 43 44}"
}

