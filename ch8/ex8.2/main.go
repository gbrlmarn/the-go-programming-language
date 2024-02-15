// Exercise 8.2: Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands from each client such as cd to change directory, ls to list a directory, get to send the contents of a file, close to close the connection. You can use the standard ftp command as the client, or write your own.
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const (
	help = "ls <path> - displays files\ncd <path> - change path\nget <file> - get file content\nclose -  closes ftp client\n"
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
		go handleConn(conn) // handle one connection at a time
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func execCmd(w io.Writer, cmdn string, args ...string) {
	cmd := exec.Command(cmdn, args...)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":
			execCmd(c, cmds[0], cmds[1:]...)
		case "cd":
			if len(cmds) < 2 {
				mustCopy(c, strings.NewReader(help))
			} else {
				if err := os.Chdir(cmds[1]); err != nil {
					log.Print(err)
				}
			}
		case "get":
			if len(cmds) < 2 {
				mustCopy(c, strings.NewReader(help))
				continue
			} else {
				f, err := os.Open(cmds[1])
				if err != nil {
					log.Printf("file %s: %v\n", cmds[1], err)
					continue
				}
				mustCopy(c, f)
			}
		case "close":
			return
		default:
			mustCopy(c, strings.NewReader(help))
		}
	}
}
