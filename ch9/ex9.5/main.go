// Exercise 9.5: Write a program with two goroutines that send messages back and forth over two unbuffered channels in pin-pong fashion. How many communications per second can the program sustain?
package main

import (
	"fmt"
	"time"
)

var shutdown = make(chan struct{})

type cch struct {
	ch chan int
	c  int
	t  time.Duration
}

func main() {
	var ping, pong cch
	ping.ch = make(chan int)
	pong.ch = make(chan int)

	go func() {
		ping.ch <- 1
		worker(&pong, &ping)
	}()
	go func() { worker(&ping, &pong) }()

	time.Sleep(3 * time.Second)
	shutdown <- struct{}{}

	fmt.Printf("%d communications in %d seconds\n", ping.c, ping.t/time.Second)
}

func worker(rcv, snd *cch) {
	start := time.Now()
	for {
		select {
		case <-shutdown:
			rcv.t = time.Since(start)
		case v := <-rcv.ch:
			rcv.c++
			snd.ch <- v
		}
	}
}
