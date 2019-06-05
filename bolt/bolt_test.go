package bolt

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestZeroChannel(t *testing.T) {
	done := make(chan struct{}, 0)
	defer close(done)
	go func() {
		for {
			select {
			case <-done:
				log.Println("=======")
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
