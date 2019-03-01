package store

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/globalsign/mgo"
)

func TestMgoSession(t *testing.T) {
	session, err := mgo.DialWithTimeout("127.0.0.1:27017", 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	// defer cancel()
	m := NewMMgo(ctx, session)
	ss := m.Session()
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				log.Println(ss.Ping())
			}
		}
	}()
	time.Sleep(2 * time.Minute)
	// log.New()
	cancel()
}
