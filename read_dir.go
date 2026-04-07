package helper

import (
	"fmt"
	"os"
)

func Readdir_directories(dirpath string) ([]string, error) {
	if len(dirpath) < 1 {
		return nil, fmt.Errorf("ERROR: Empty directory path.")
	}

	// ReadDir reads the named directory and returns
	// a list of directory entries sorted by filename.
	directories, err := os.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	if (directories == nil) || (len(directories) < 1) {
		return nil, fmt.Errorf("ERROR: Empty directory.")
	}

	dircnt := len(directories)
	dirlist := make([]string, dircnt)
	for i, dir := range files {
		if !file.IsDir() {
			continue
		}
		dirlist[i] = dir
	}

	return dirlist, nil
}


func Readdir_files(dirpath string) ([]string, error) {
	if len(dirpath) < 1 {
		return nil, fmt.Errorf("ERROR: Empty directory path.")
	}

	// ReadDir reads the named directory and returns
	// a list of directory entries sorted by filename.
	files, err := os.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	if (files == nil) || (len(files) < 1) {
		return nil, fmt.Errorf("ERROR: Empty directory.")
	}

	filecnt := len(files)
	/* for _, file := range files {
		// Use file.Name() to get the string name
		// Use file.IsDir() if you need to filter out directories.
		fmt.Println(file.Name())
		if !file.IsDir() {
			filecnt++
		}
	} */

	filelist := make([]string, filecnt)
	for i, file := range files {
		if file.IsDir() {
			continue
		}
		filelist[i] = file
	}

	return filelist, nil
}
