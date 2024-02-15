package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x, string(b.op), b.y)
}

func (c call) String() string {
	b := &bytes.Buffer{}
	b.WriteString(c.fn)
	b.WriteString("(")
	for i, arg := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteString(")")
	return b.String()
}

func (m min) String() string {
	b := &bytes.Buffer{}
	b.WriteString("[")
	for i, arg := range m.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteString("]")
	return b.String()
}
