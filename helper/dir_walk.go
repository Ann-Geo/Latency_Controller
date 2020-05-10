package helper

import (
	"os"
	"path/filepath"
)

func WalkAllFilesInDir(dir string) (string, []string) {
	var errMsg string = "file read success"
	fileList := make([]string, 0)
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errMsg = "Incorrect file path"
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			fileList = append(fileList, path)

		}
		return nil
	})

	return errMsg, fileList
}
