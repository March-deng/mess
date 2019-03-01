package main

import (
	"log"
	"os"
	"strconv"
	"time"
)


var openFile os.File
var totalWrites int
var entryPerLogFile int = 10
var whenToStop int = 23

func main() {
	log.Println("test")
	numberOfLogFiles := 0

	filename := "/tmp/mylog.log"

	err := setUpLogFile(filename)
	log.Println(err)
	if err != nil {
		log.Println("initial error")
		os.Exit(-1)
	}
	// fmt.Println("start")
	for {
		log.Println(numberOfLogFiles, "this is a testing use of a log file")
		numberOfLogFiles++
		totalWrites++
		// fmt.Println(numberOfLogFiles, totalWrites)
		if numberOfLogFiles >= entryPerLogFile {
			rotateFiles(filename)
			numberOfLogFiles = 0
		}
		if totalWrites > whenToStop {
			// fmt.Println("rotate")
			rotateFiles(filename)
			break
		}
		time.Sleep(1 * time.Millisecond)

	}
	log.Println("end and exit")
}

func rotateFiles(filename string) error {
	openFile.Close()
	os.Rename(filename, filename+"."+strconv.Itoa(totalWrites))
	return setUpLogFile(filename)
}

func setUpLogFile(filename string) error {
	openFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(openFile)
	return nil
}
