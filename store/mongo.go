package store

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"github.com/globalsign/mgo"
)

type MMgo struct {
	v      atomic.Value
	exp    time.Time
	ticker *time.Ticker
}

const (
	checkTimeInterval = time.Minute
)

func NewMMgo(ctx context.Context, session *mgo.Session) *MMgo {
	m := &MMgo{}
	ss := session.Copy()
	m.v.Store(ss)
	m.ticker = time.NewTicker(checkTimeInterval)
	m.exp = time.Now().Add(checkTimeInterval)
	go m.refresh(ctx)
	return m
}

func (m *MMgo) Session() *mgo.Session {
	ss := m.v.Load().(*mgo.Session)
	if time.Now().After(m.exp) {
		ss.Refresh()
		m.exp = time.Now().Add(checkTimeInterval)
		m.v.Store(ss)
	}
	return ss
}

func (m *MMgo) refresh(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("close the refresh routine with:", ctx.Err())
			m.ticker.Stop()
			return
		case <-m.ticker.C:
			m.Session()
		}
	}
}
