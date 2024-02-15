// Exercise 10.4: Construct a tool that reports the set of all packages in the workspace that transitively depend on the packages specified by the arguments. Hint: you will need to run go list twice, once for the initial packages and once for all packages. You may want to parse its JSON output using the encoding/json package (~4.5).
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
)

type Pkg struct {
	ImpPath string   `json:"ImportPath"`
	Name    string   `json:"Name"`
	Deps    []string `json:"Deps"`
}

func main() {
	key := os.Args[1]
	cmd := exec.Command("go", "list", "-json", "...")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	var s []byte
	var buf bytes.Buffer
	for _, b := range out {
		switch b {
		case '{':
			s = append(s, b)
		case '}':
			s = s[0 : len(s)-1]
		}
		buf.WriteByte(b)
		if b == '}' && len(s) == 0 {
			var p Pkg
			err := json.Unmarshal(buf.Bytes(), &p)
			if err != nil {
				log.Fatal(err)
			}
			if sort.SearchStrings(p.Deps, key) != len(p.Deps) {
				fmt.Println(p.ImpPath)
			}
			buf.Truncate(0)
		}
	}
}
