// Exercise 13.4: Depending on C libraries has its drawbacks. Provide an alternative pure-Go implementation of bzip.NewWriter that uses the os/exec package to run /bin/bzip2 as a subprocess.

// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip 

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	wc   io.WriteCloser // underlying output stream
	mu  sync.Mutex
	cmd *exec.Cmd
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("/bin/bzip2")
	cmd.Stdout = out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	w := &writer{wc: stdin, cmd: cmd}
    err = cmd.Start()
    if err != nil {
        return nil, err
    }
    return w, nil
}

func (w *writer) Write(data []byte) (int ,error) {
    w.mu.Lock()
    defer w.mu.Unlock()
    return w.wc.Write(data)
}

func (w *writer) Close() error {
    w.mu.Lock()
    defer w.mu.Unlock()
    if err := w.wc.Close() ; err != nil {
        return err
    }
    if err := w.cmd.Wait(); err != nil {
        return err
    }
    return nil
}
