// Package webdavfs implements webdav.FileSystem using an http.FileSystem.
package webdavfs

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

// New returns a webdav.FileSystem adapter for the provided http.FileSystem.
func New(fs http.FileSystem) webdav.FileSystem {
	return &webdavFS{fs: fs}
}

type webdavFS struct {
	fs http.FileSystem
}

func (w *webdavFS) OpenFile(_ context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	if flag != os.O_RDONLY {
		return nil, os.ErrPermission
	}
	f, err := w.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return &webdavFile{File: f}, nil
}

func (w *webdavFS) Stat(_ context.Context, name string) (os.FileInfo, error) {
	f, err := w.fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Stat()
}

func (*webdavFS) Mkdir(_ context.Context, name string, perm os.FileMode) error {
	return os.ErrPermission
}

func (*webdavFS) RemoveAll(_ context.Context, name string) error {
	return os.ErrPermission
}

func (*webdavFS) Rename(_ context.Context, oldName, newName string) error {
	return os.ErrPermission
}

type webdavFile struct {
	http.File
}

func (*webdavFile) Write([]byte) (int, error) {
	return 0, &os.PathError{Op: "write", Err: os.ErrInvalid}
}
