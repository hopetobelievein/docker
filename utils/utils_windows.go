// +build daemon,windows

package utils

import (
	"os"
	"path/filepath"
)

// TreeSize walks a directory tree and returns its total size in bytes.
func TreeSize(dir string) (size int64, err error) {
	err = filepath.Walk(dir, func(d string, fileInfo os.FileInfo, e error) error {
			// Ignore directory sizes
			if fileInfo == nil {
				return nil
			}

			s := fileInfo.Size()
			if fileInfo.IsDir() || s == 0 {
				return nil
			}

			size += s

			return nil
		})
	return
}
