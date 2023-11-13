package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, host string
}

func (c clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}

func main() {
	clocks := parse(os.Args[1:])
	for _, c := range clocks {
		go listen(c)
	}
	
	for {
		time.Sleep(time.Minute)
	}
}

func parse(args []string) []clock {
	var clocks []clock
	for _, a := range args {
		tmp := strings.Split(a, "=")
		if len(tmp) != 2 {
			log.Fatal("parse: name=hostname:port\n")
		}
		clocks = append(clocks, clock{tmp[0], tmp[1]})
	}
	return clocks
}

func listen(c clock) {
	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c.watch(os.Stdout, conn)
}

