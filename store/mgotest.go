package store

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/dbtest"
)

type TestUtil struct {
	started bool
	dbPath  string
	server  dbtest.DBServer
	ss      *mgo.Session
}

func (tu *TestUtil) Session(t *testing.T) *mgo.Session {
	if !tu.started {
		dbPath := filepath.Join(os.TempDir(), "mgotest", strconv.FormatInt(time.Now().UnixNano(), 16))
		err := os.MkdirAll(dbPath, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		tu.dbPath = dbPath
		tu.server.SetPath(dbPath)
		tu.started = true
		tu.ss = tu.server.Session()
	}
	return tu.ss
}

func (tu *TestUtil) Close() {
	if tu.started {
		tu.ss.Close()
	}
	tu.server.Stop()
	tu.server.Wipe()
	if tu.started {
		os.RemoveAll(tu.dbPath)
	}
}
