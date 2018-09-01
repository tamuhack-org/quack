package database

import (
	"github.com/globalsign/mgo"
	"time"
	// Reading in environmental variables.
	"github.com/tamuhack-org/quack/backend/shared/config"
)

// This is a persistent session that's accessed by our handlers.
var (
	Mongo *mgo.Session
)

// [Constructor] Connect to the database
func CreateAndConnect() error {
	var err error

	// Global config from env vars.
	globalConfig := config.GlobalConfig

	// Dial connects to the remote db.
	if Mongo, err = mgo.DialWithTimeout(globalConfig.MongoUrl, 5*time.Second); err != nil {
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
