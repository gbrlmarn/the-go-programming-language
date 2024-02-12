// Exercise 12.1: Extend Display so that it can display maps whose keys are struct or arrays.
package display

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func formatAtom(v reflect.Value) string {
    switch v.Kind() {
    case reflect.Invalid:
        return "invalid"
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return strconv.FormatInt(v.Int(), 10)
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return strconv.FormatUint(v.Uint(), 10)
    // ...floating-point and complex cases omitted for brevity... 
    case reflect.Bool:
        if v.Bool() {
            return "true"
        }
        return "false"
    case reflect.String:
        return strconv.Quote(v.String())
    case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
        return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
    default: // reflect.Array, reflect.Struct, reflect.Interface
        return v.Type().String() + " value"
    }
}

func formatKey(v reflect.Value) string {
    switch v.Kind() {
    case reflect.Struct:
        b := &bytes.Buffer{} 
        b.WriteString("{")
        for i := 0; i < v.NumField(); i++ {
            if i != 0 {
                b.WriteString(", ")
            }
            structEl := fmt.Sprintf("{%s: %s}",v.Type().Field(i).Name, formatAtom(v.Field(i)))
            b.WriteString(structEl)
        }
        b.WriteString("}")
        return b.String() 
    case reflect.Array:
        b := &bytes.Buffer{} 
        b.WriteString("{")
        for i := 0; i < v.Len(); i++ {
            if i != 0 {
                b.WriteString(", ")
            }
            b.WriteString(formatAtom(v.Index(i)))
        }
        b.WriteString("}")
        return b.String() 
    default:
        return formatAtom(v)
    }
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
            display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
    case reflect.Struct:
        for i := 0; i < v.NumField(); i++ {
            fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
            display(fieldPath, v.Field(i))
        }
    case reflect.Map:
        for _, key := range v.MapKeys() {
            display(fmt.Sprintf("%s[%s]", path, formatKey(key)), v.MapIndex(key))
        }
    case reflect.Ptr:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            display(fmt.Sprintf("(*%s)", path), v.Elem())
        }
    case reflect.Interface:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
            display(path+".value", v.Elem())
        }
    default: // basic types, channels, funcs
        fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}