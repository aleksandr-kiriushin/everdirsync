package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listDir(directory string) error {
	f, err := os.Open(directory)
	if err != nil {
		return err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return err
	}

	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			err := listDir(filepath.Join(directory, name))
			if err != nil {
				return err
			}
		} else {
			fmt.Println("File:", filepath.Join(directory, name))
		}
	}
	return nil
}
