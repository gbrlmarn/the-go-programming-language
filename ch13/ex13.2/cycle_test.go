package cycle

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	one, oneAgain := 1, 1

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1 := make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1, iface1Again interface{} = &one, &oneAgain

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{"foo", false},
		{"foo", false},
		{mystring("foo"), false}, // different types
		// slices
		{[]string{"foo"}, false},
		{[]string{"foo"}, false},
		{[]string{}, false},
		// slice cycles
		{cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{},
			false,
		},
		// pointers
		{&one, false},
		{&one, false},
		{&one, false},
		{new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
		{cyclePtr1, true}, // they're deeply equal
		// functions
		{(func())(nil), false},
		{(func())(nil), false},
		{func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, false},
		{[...]int{1, 2, 3}, false},
		// channels
		{ch1, false},
		{ch1, false},
		{ch1ro, false}, // NOTE: not equal
		// interfaces
		{&iface1, false},
		{&iface1, false},
		{&iface1Again, false},
	} {
		if Cycle(test.x) != test.want {
			t.Errorf("Cycle(%v) = %t", test.x, !test.want)
		}
	}
}

func Example_cycle1() {
	fmt.Println(Cycle([]string{"foo"}))     // "false"
	fmt.Println(Cycle([]string(nil)))       // "true"
	fmt.Println(Cycle(map[string]int(nil))) // "true"

	// Output:
	// false
	// false 
	// false 
}

func Example_cycle2() {
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Cycle(a)) // "true"
	fmt.Println(Cycle(b)) // "true"
	fmt.Println(Cycle(c)) // "true"
	fmt.Println(Cycle(a)) // "true"
	fmt.Println(Cycle(a)) // "true"

	// Output:
	// true
	// true
	// true
	// true 
	// true 
}
