// Exercise 12.5: Adapt encode to emit JSON instead of S-expression. Test your encoder using the standard decoder, json.Unmarshal.
package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
    var buf bytes.Buffer
    if err := encode(&buf, reflect.ValueOf(v)); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
    switch v.Kind() {
    case reflect.Invalid:
        buf.WriteString("null")
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        fmt.Fprintf(buf, "%d", v.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        fmt.Fprintf(buf, "%d", v.Uint())
    case reflect.String:
        fmt.Fprintf(buf, "%q", v.String())
    case reflect.Ptr:
        return encode(buf, v.Elem())
    case reflect.Array, reflect.Slice: // (value ...)
        buf.WriteByte('[')
        for i := 0; i < v.Len(); i++ {
            if i > 0 {
                buf.WriteByte(',')
            }
            if err := encode(buf, v.Index(i)); err != nil {
                return err
            }
        }
        buf.WriteByte(']')
    case reflect.Struct: // ((name value) ...)
        buf.WriteByte('{')
        for i := 0; i < v.NumField(); i++ {
            if i > 0 {
                buf.WriteByte(',')
            }
            fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
            if err := encode(buf, v.Field(i)); err != nil {
                return err
            }
        } 
        buf.WriteByte('}')
    case reflect.Map: // ((key value) ...)
        buf.WriteByte('{')
        for i, key := range v.MapKeys() {
            if i > 0 {
                buf.WriteByte(',')
            }
            if err := encode(buf, key); err != nil {
                return err
            }
            buf.WriteByte(':')
            if err := encode(buf, v.MapIndex(key)); err != nil {
                return err
            }
        }
        buf.WriteByte('}')
    default: // float, complex, bool, chan, func, interface
        return fmt.Errorf("unsupported type: %s", v.Type())
    }
    return nil
}
