// Exercise 8.4: Modify the reverb2 server to use a sync.WaitGroup per connection to count the number of active echo goroutines. When it falls to zero, close the write half of the TCP connection as described in Exercise 8.3. Verify that your modified netcat3 client from the exercise waits for t he final echoes of multiple concurrent shouts, even after the standard input has been closed.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Print(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup // number of working goroutines
	for input.Scan() {
		wg.Add(1)
		// worker
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}

	// closer
	// NOTE: ignoring potential errors from input.Err()
	go func() {
		wg.Wait()
		if ctcp, ok := c.(*net.TCPConn); ok {
			ctcp.CloseWrite()
		} else {
			c.Close()
		}
	}()
}
