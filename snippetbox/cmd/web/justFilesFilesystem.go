package main

import (
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	Fs http.FileSystem
}

/* Disable directory path */
func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.Fs.Open(name)
	
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}