package main

import (
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	s := []string{
		"uno", "uno", "due",
		"due", "due", "due",
		"tre", "tre", "quattro",
		"quattro", "quattro",
		"cinque", "cinque",
	}
	got := nodup(s)
	want := []string{
		"uno", "due", "tre",
		"quattro", "cinque",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n", got, want)
	}
}
