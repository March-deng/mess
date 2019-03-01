package store

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"github.com/globalsign/mgo"
)

type Mgo struct {
	session   atomic.Value
	dburl     string
	sleepTime <-chan time.Time
}

var nullSession *mgo.Session

const (
	defaultTimeOut = 10 * time.Second

	failTimeOut = 10 * time.Second

	pingTimeOut = 3 * time.Second
)

func NewMgo(ctx context.Context, dburl string) *Mgo {
	log.Println("connect to the mongodb at:", dburl)
	m := &Mgo{dburl: dburl}
	m.session.Store(nullSession)
	m.checkSession(ctx)
	return m
}

func (m *Mgo) Session() *mgo.Session {
	return m.session.Load().(*mgo.Session)
}

func (m *Mgo) Ping() error {
	return m.Session().Ping()
}

//checkSession method keep alive the mgo session if it is available
func (m *Mgo) checkSession(ctx context.Context) {
	session, err := mgo.DialWithTimeout(m.dburl, defaultTimeOut)
	if err != nil {
		log.Println("CAUTION: check if the mongodb in work.")
		m.sleepTime = time.After(failTimeOut)
	}
	m.session.Store(session)
	m.sleepTime = time.After(pingTimeOut)
	for {
		select {
		case <-m.sleepTime:
			err = m.Session().Ping()
			if err != nil {
				log.Println("WARNING: mongodb connection invalid, please check.")
				m.sleepTime = time.After(pingTimeOut)
				m.session.Store(nullSession)
			}
		case <-ctx.Done():
			log.Println("exit the mgo session keep alive loop because of the context exit.")
			return
		}
	}
}
