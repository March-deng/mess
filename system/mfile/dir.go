package mfile

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//copy dir copy the path dir to a new path,
//but do not copy any existing files.
func CopyDir(current, new string) error {
	permissions := os.ModePerm
	_, err := os.Stat(new)
	if os.IsNotExist(err) {
		os.MkdirAll(new, permissions)
	} else {
		log.Printf("the new path: %s already exists, exit...", new)
		os.Exit(1)
	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		fileInfo, _ := os.Lstat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			log.Println("skipping", path)
			return nil
		}
		fileInfo, err = os.Stat(path)
		if err != nil {
			log.Println("*", err)
			return err
		}

		mode := fileInfo.Mode()
		if mode.IsDir() {
			tempPath := strings.Replace(path, current, "", 1)
			pathToCreate := new + "/" + filepath.Base(current) + tempPath
			log.Println(":", pathToCreate)
			_, err := os.Stat(pathToCreate)
			if os.IsNotExist(err) {
				os.MkdirAll(pathToCreate, permissions)
			} else {
				log.Println("did not create", pathToCreate, ":")
			}
		}
		return nil
	}
	return filepath.Walk(current, walkFunc)
}

func copy() {
	io.Copy()
}
