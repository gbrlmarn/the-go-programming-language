// Exercise 13.2: Write a function that reports whether its argument is a cyclic data structure.
package cycle

import (
	"reflect"
	"unsafe"
)

func cycle(v reflect.Value, seen map[comparison]bool) bool {
	// cycle check
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		c := comparison{vptr, v.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cycle(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if cycle(v.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if cycle(v.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range v.MapKeys() {
			if cycle(v.MapIndex(k), seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func Cycle(v any) bool {
	seen := make(map[comparison]bool)
	return cycle(reflect.ValueOf(v), seen)
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}
