// Exercise 7.18: Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it. Nodes are of two kinds: CharData nodes represent text strings, and Element nodes represent named elements and their attributes. Each element node has a slice of child nodes.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Node interface{
	String() string
} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (c CharData) String() string {
	return string(c)
}

func (e Element) String() string {
	var attrs, children string
	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}
	for _, child := range e.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s%s>%s</%s>", e.Type.Local, attrs, children, e.Type.Local)
}

func visit(n Node, d int, w io.Writer) {
	switch n := n.(type) {
	case Element:
		fmt.Fprintf(w, "%*s%s %s\n", d*2, "", n.Type.Local, n.Attr)
		for _, c := range n.Children {
			visit(c, d+1, w)
		}
	case CharData:
		fmt.Fprintf(w, "%*s%q\n", d*2, "", n)
	}
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []Element // stack of elements
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			el := Element{tok.Name, tok.Attr, nil}
			if len(stack) > 0  {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, el)
			}
			stack = append(stack, el) // push
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) > 0  {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, CharData(tok))
			}
		}
	}
}

func main() {
	xmlf, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	root, err := parse(xmlf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(root)
}
 
