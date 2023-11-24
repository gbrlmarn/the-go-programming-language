// Exercise 8.9: Write a version of du that computes and periodically displays separate totals for each of the root directories.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

var done = make(chan struct{})

type root struct {
	i int
	s int64
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	rootSizes := make(chan root)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, i, rootSizes)
	}
	go func() {
		n.Wait()
		close(rootSizes)		
	}()
	
	// Print the results periodically.
	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case rsize, ok := <- rootSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[rsize.i]++
			nbytes[rsize.i] += rsize.s
		case <- tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%s:\t%d files, %.1f GB\n", r, nfiles[i], float64(nbytes[i])/1e9)
	}
}

// walkDir recursively walks the files tree rooted at dir
// and sends the size of each found file on filesSizes.
func walkDir(dir string, n *sync.WaitGroup, i int, rootSizes chan<- root) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, i, rootSizes)
		} else {
			inf, err := entry.Info()
			if err != nil {
				log.Print(err)
			}
			rootSizes <- root{i, inf.Size()}
		}
	}
}

// sema is a counting semaphore for limiting councurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents return the entries of directory dir.
func dirents(dir string) []fs.DirEntry {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token
	
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

