// Exercise 8.8: Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any client that shouts nothing withing 10 seconds.
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


func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	text := make(chan string)
	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)
	
	var wg sync.WaitGroup
	defer func() {
		c.Close()
		wg.Wait()
	}()

	go func() {
		for input.Scan() {
			text <- input.Text()
		}
		close(text)
	}()
	
	for {
		select {
		case t := <-text:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, t, 1*time.Second, &wg)
		case <-timer.C:
			fmt.Fprintln(c, "Times up")
			return
		}
	}
}


