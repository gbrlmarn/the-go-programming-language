// Exercise 11.2: Write a set of tests for IntSet (§6.5) that checks that its behaviour after each operation is equivalent to a set based on built-in maps. Save your implementation for benchmarking in Exercise 11.7.
package main

import (
	"gopl/ch11/ex11.2/hashset"
	"gopl/ch11/ex11.2/intset"
	"testing"
)

func TestIntSetHashSet(t *testing.T) {
    iset1 := intset.IntSet{}
    hset1 := make(hashset.HashSet)
    iset1.Add(1)
    iset1.Add(2)
    hset1.Add(1)
    hset1.Add(2)
    if hset1.String() != iset1.String() {
        t.Errorf("%v != %v", hset1.String(), iset1.String())
    }

    iset2 := intset.IntSet{}
    hset2 := make(hashset.HashSet)
    iset2.Add(3)
    iset2.Add(4)
    hset2.Add(3)
    hset2.Add(4)
    iset1.UnionWith(&iset2)
    hset1.UnionWith(hset2)
    if hset1.String() != iset1.String() {
        t.Errorf("%v != %v", hset1.String(), iset1.String())
    }
}
