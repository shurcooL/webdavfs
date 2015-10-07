// Package vfsutil implements some I/O utility functions for webdav.FileSystem.
package vfsutil

import (
	"io"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

// Open opens the named file for reading.  If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(fs webdav.FileSystem, name string) (http.File, error) {
	return fs.OpenFile(name, os.O_RDONLY, 0)
}

// WriteFile writes data to a file named by name.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing.
func WriteFile(fs webdav.FileSystem, name string, data []byte, perm os.FileMode) error {
	f, err := fs.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
