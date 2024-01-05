package zip

import (
	"archive/zip"
	"fmt"
	"gopl/ch10/ex10.2/arch"
)

func Open(fileName string) error {
    zr, err := zip.OpenReader(fileName)
    if err != nil {
        return err
    }
    defer zr.Close()
    for _, f := range zr.File {
        fmt.Printf("%s\n", f.Name)
    }
    return nil
}

func init() {
    arch.Register("zip", Open)
}
