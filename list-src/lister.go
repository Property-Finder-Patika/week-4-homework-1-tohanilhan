package main

import (
	"io/fs"
	"path/filepath"
)

var (
	path      = ".."  // path to the directory to be scanned
	extension = ".go" // extension of the files to be scanned
)

func main() {

	filepath.WalkDir(path, func(fileName string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// get only go source codes from the directory

		// if not a directory, then check if the file extension is .go
		if !dir.IsDir() && filepath.Ext(dir.Name()) == extension {
			println(fileName)
		}
		// call WalkDir again if current file is a directory
		return nil
	})

	// or you can use it like below, to see, you can comment the above code and use this one
	// filepath.WalkDir(path, WalkDirectory)

}

func WalkDirectory(fileName string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	// get only go source codes from the directory
	if !d.IsDir() && filepath.Ext(d.Name()) == extension {
		println(fileName)
	}

	return nil
}
