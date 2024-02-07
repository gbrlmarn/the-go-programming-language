// Exercise 12.7: Create a streaming API for S-expression encoder, following the style of json.Encoder (§4.5).

package sexpr

import (
	"bytes"
	"fmt"
	"io"
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
        buf.WriteString("nil")
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        fmt.Fprintf(buf, "%d", v.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        fmt.Fprintf(buf, "%d", v.Uint())
    case reflect.String:
        fmt.Fprintf(buf, "%q", v.String())
    case reflect.Ptr:
        return encode(buf, v.Elem())
    case reflect.Array, reflect.Slice: // (value ...)
        buf.WriteByte('(')
        for i := 0; i < v.Len(); i++ {
            if i > 0 {
                buf.WriteByte(' ')
            }
            if err := encode(buf, v.Index(i)); err != nil {
                return err
            }
        }
        buf.WriteByte(')')
    case reflect.Struct: // ((name value) ...)
        buf.WriteByte('(')
        for i := 0; i < v.NumField(); i++ {
            if i > 0 {
                buf.WriteByte(' ')
            }
            fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
            if err := encode(buf, v.Field(i)); err != nil {
                return err
            }
            buf.WriteByte(')')
        } 
        buf.WriteByte(')')
    case reflect.Map: // ((key value) ...)
        buf.WriteByte('(')
        for i, key := range v.MapKeys() {
            if i > 0 {
                buf.WriteByte(' ')
            }
            buf.WriteByte('(')
            if err := encode(buf, key); err != nil {
                return err
            }
            buf.WriteByte(' ')
            if err := encode(buf, v.MapIndex(key)); err != nil {
                return err
            }
            buf.WriteByte(')')
        }
        buf.WriteByte(')')
    case reflect.Bool:
        if v.Bool() {
            buf.WriteByte('t')
        } else {
            buf.WriteString("nil")
        }
    case reflect.Float32, reflect.Float64:
        fmt.Fprintf(buf, "%g", v.Float())
    case reflect.Complex64, reflect.Complex128:
        fmt.Fprintf(buf, "#C(%g %g)", real(v.Complex()), imag(v.Complex()))
    case reflect.Interface:
        var b bytes.Buffer
        encode(&b, v.Elem())
        fmt.Fprintf(buf, "(%q %s)", v.Elem().Type(), b.String())
    default: // chan, func 
        return fmt.Errorf("unsupported type: %s", v.Type())
    }
    return nil
}

// An Encoder writes S-expr values to an output stream.
type Encoder struct {
    w io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
    return &Encoder{w: w}
}

// Encode writes the S-expr encoding of v to the stream,
func (enc *Encoder) Encode(v any) error {
    b, err := Marshal(v)
    if err != nil {
        return err
    }
    _, err = enc.w.Write(b)
    if err != nil {
        return err
    }
    return nil
}

