package bolt

import (
	"github.com/boltdb/bolt"
	log "github.com/sirupsen/logrus"
)

func openBolt() {
	db, err := bolt.Open("my.db", 0777, nil)
	if err != nil {
		log.Println("open bolt db error")
		return
	}
	defer db.Close()

}
