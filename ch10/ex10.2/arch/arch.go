// Exercise 10.2: Define a generic archive file-reading function capable of reading ZIP files(archive/zip) and POSIX tar files (archive/tar). Use a registration mechanism similar to the one described above so that support for each file format can be plugged in using blank imports.
package arch

import (
	"fmt"
	"path/filepath"
)

type format struct {
	name   string
	reader NewReader
}

type NewReader func(fileName string) error

var formats []format

func Register(n string, f func(fileName string) error) {
	formats = append(formats, format{name: n, reader: f})
}

func Open(fileName string) error {
	formatName := filepath.Ext(fileName)
	formatName = formatName[1:]
	for _, format := range formats {
		if formatName == format.name {
			return format.reader(fileName)
		}
	}
	return fmt.Errorf("archive format not found: %v\n", formatName)
}
