package database

import (
	"github.com/globalsign/mgo"
	"time"
)

// This is a persistent session that's accessed by our handlers.
var (
	Mongo *mgo.Session
)

// [Constructor] Connect to the database
func CreateAndConnect(url string) error {
	var err error
	// Dial connects to the remote db.
	if Mongo, err = mgo.DialWithTimeout(url, 5*time.Second); err != nil {
		return err
	}

	// Set timeout and check that its alive.
	Mongo.SetSocketTimeout(1 * time.Second)
	if err = Mongo.Ping(); err != nil {
		return err
	}
	if err = Mongo.Ping(); err != nil {
		return err
	}

  return nil
}
