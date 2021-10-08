package goscriptable

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func cloneZipItem(f *zip.File, dest string) bool {
	// Create full directory path
	path := filepath.Join(dest, f.Name)
	err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
	if err != nil {
		return false
	}

	// Clone if item is a file
	rc, err := f.Open()
	if err != nil {
		return false
	}
	if !f.FileInfo().IsDir() {
		fileCopy, err := os.Create(path)
		if err != nil {
			return false
		}
		_, err = io.Copy(fileCopy, rc)
		fileCopy.Close()
		if err != nil {
			return false
		}
	}
	rc.Close()
	return true
}
