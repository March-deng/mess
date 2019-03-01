package core

import (
	"log"
	"testing"
)

func TestGetData(t *testing.T) {
	p := GetFileData()
	log.Println(p)
}

func TestChan(t *testing.T) {
	sig := make(chan string)
	go func() {
		sig <- "dengcong"
	}()
	log.Println(<-sig)
}
