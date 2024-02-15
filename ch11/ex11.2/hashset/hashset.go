package hashset

import (
	"bytes"
	"fmt"
	"sort"
)

// An HashSet is a set of small non-negative integers.
type HashSet map[int]bool

// Has reports if the set contains x value
func (s HashSet) Has(x int) bool {
	return s[x]
}

// Add adds the non-negative value x to the set.
func (s HashSet) Add(x int) {
	s[x] = true
}

// UnionWith sets s to the union of s and t.
func (s HashSet) UnionWith(t HashSet) {
	for k := range t {
		s[k] = true
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s HashSet) String() string {
	var ks []int
	for k := range s {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range ks {
		fmt.Fprintf(&buf, "%d", v)
		if i < len(ks)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
