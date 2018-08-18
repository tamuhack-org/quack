package database

import (
	"github.com/globalsign/mgo"
	"time"
)

// Database object holds a pointer to a session.
type Database struct {
	// Database type
	session *mgo.Session
}

// [Constructor] Connect to the database
func CreateAndConnect(url string) (*Database, error) {
	// Dial connects to the remote db.
	db, err := mgo.DialWithTimeout(url, 5*time.Second)
	if err != nil {
		return nil, err
	}

	// Set timeout and check that its alive.
	db.SetSocketTimeout(1 * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	obj := Database{session: db}
	return &obj, nil
}
