// Exercise 11.7: Write benchmarks for Add, UnionWith, and other methods of *IntSet(𝕊6.5) using large pseudo-random inputs. How fast can you make these methods run? How does the choice of words size affect performance? How fast is IntSet compared to a set implementation based of the built-in map type?
package main

import (
	"gopl/ch11/ex11.2/hashset"
	"gopl/ch11/ex11.2/intset"
	"math"
	"math/rand"
	"testing"
	"time"
)

func initIntSet() intset.IntSet {
	return intset.IntSet{}
}

func addAllIntSet(s intset.IntSet, ints []int) intset.IntSet {
	for i := 0; i < len(ints); i++ {
		s.Add(ints[i])
	}
	return s
}

func initHashSet() hashset.HashSet {
	return make(hashset.HashSet)
}

func addAllHashSet(s hashset.HashSet, ints []int) hashset.HashSet {
	for i := 0; i < len(ints); i++ {
		s.Add(ints[i])
	}
	return s
}

func randInts(n int) []int {
	ints := make([]int, 0, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		ints = append(ints, rand.Intn(math.MaxInt8))
	}
	return ints
}

func benchmarkAdd(b *testing.B, n int, f func(int)) {
	ints := randInts(n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(ints[j])
		}
	}
}

func benchmarkUnionIntSets(b *testing.B, n int) {
	ints1 := randInts(n)
	ints2 := randInts(n)
	s1 := initIntSet()
	s2 := initIntSet()
	addAllIntSet(s1, ints1)
	addAllIntSet(s2, ints2)
	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}

func benchmarkUnionHashSets(b *testing.B, n int) {
	ints1 := randInts(n)
	ints2 := randInts(n)
	s1 := initHashSet()
	s2 := initHashSet()
	addAllHashSet(s1, ints1)
	addAllHashSet(s2, ints2)
	for i := 0; i < b.N; i++ {
		s1.UnionWith(s2)
	}
}

func BenchmarkAddIntSet1(b *testing.B) {
	s := initIntSet()
	benchmarkAdd(b, 1, s.Add)
}
func BenchmarkAddIntSet10(b *testing.B) {
	s := initIntSet()
	benchmarkAdd(b, 10, s.Add)
}
func BenchmarkAddIntSet100(b *testing.B) {
	s := initIntSet()
	benchmarkAdd(b, 100, s.Add)
}
func BenchmarkAddIntSet1000(b *testing.B) {
	s := initIntSet()
	benchmarkAdd(b, 1000, s.Add)
}
func BenchmarkAddIntSet10000(b *testing.B) {
	s := initIntSet()
	benchmarkAdd(b, 10000, s.Add)
}

func BenchmarkAddHashSet1(b *testing.B) {
	s := initHashSet()
	benchmarkAdd(b, 1, s.Add)
}
func BenchmarkAddHashSet10(b *testing.B) {
	s := initHashSet()
	benchmarkAdd(b, 10, s.Add)
}
func BenchmarkAddHashSet100(b *testing.B) {
	s := initHashSet()
	benchmarkAdd(b, 100, s.Add)
}
func BenchmarkAddHashSet1000(b *testing.B) {
	s := initHashSet()
	benchmarkAdd(b, 1000, s.Add)
}
func BenchmarkAddHashSet10000(b *testing.B) {
	s := initHashSet()
	benchmarkAdd(b, 10000, s.Add)
}

func BenchmarkUnionIntSet1(b *testing.B) {
	benchmarkUnionIntSets(b, 1)
}
func BenchmarkUnionIntSet10(b *testing.B) {
	benchmarkUnionIntSets(b, 10)
}
func BenchmarkUnionIntSet100(b *testing.B) {
	benchmarkUnionIntSets(b, 100)
}
func BenchmarkUnionIntSet1000(b *testing.B) {
	benchmarkUnionIntSets(b, 1000)
}
func BenchmarkUnionIntSet10000(b *testing.B) {
	benchmarkUnionIntSets(b, 10000)
}

func BenchmarkUnionHashSet1(b *testing.B) {
	benchmarkUnionHashSets(b, 1)
}
func BenchmarkUnionHashSet10(b *testing.B) {
	benchmarkUnionHashSets(b, 10)
}
func BenchmarkUnionHashSet100(b *testing.B) {
	benchmarkUnionHashSets(b, 100)
}
func BenchmarkUnionHashSet1000(b *testing.B) {
	benchmarkUnionHashSets(b, 1000)
}
func BenchmarkUnionHashSet10000(b *testing.B) {
	benchmarkUnionHashSets(b, 10000)
}

// goos: linux
// goarch: amd64
// pkg: gopl/ch11/ex11.7
// cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
// BenchmarkAddIntSet1-6          	270497084	         4.449 ns/op
// BenchmarkAddIntSet10-6         	30182679	        39.74 ns/op
// BenchmarkAddIntSet100-6        	 3017992	       397.9 ns/op
// BenchmarkAddIntSet1000-6       	  304719	      3931 ns/op
// BenchmarkAddIntSet10000-6      	   30627	     39232 ns/op
// BenchmarkAddHashSet1-6         	135626444	         8.834 ns/op
// BenchmarkAddHashSet10-6        	11042378	       110.6 ns/op
// BenchmarkAddHashSet100-6       	 1000000	      1215 ns/op
// BenchmarkAddHashSet1000-6      	   65316	     19073 ns/op
// BenchmarkAddHashSet10000-6     	    5562	    219450 ns/op
// BenchmarkUnionIntSet1-6        	1000000000	         0.7826 ns/op
// BenchmarkUnionIntSet10-6       	1000000000	         0.7857 ns/op
// BenchmarkUnionIntSet100-6      	1000000000	         0.7831 ns/op
// BenchmarkUnionIntSet1000-6     	1000000000	         0.7849 ns/op
// BenchmarkUnionIntSet10000-6    	1000000000	         0.7850 ns/op
// BenchmarkUnionHashSet1-6       	29430301	        40.90 ns/op
// BenchmarkUnionHashSet10-6      	 5306268	       236.4 ns/op
// BenchmarkUnionHashSet100-6     	  774777	      1477 ns/op
// BenchmarkUnionHashSet1000-6    	  345552	      3416 ns/op
// BenchmarkUnionHashSet10000-6   	  339375	      3386 ns/op
// PASS
// ok  	gopl/ch11/ex11.7	25.298s
