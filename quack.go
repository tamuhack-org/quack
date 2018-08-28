package main

import (
	"log"
	"runtime"

	// Routes.
	"github.com/tamuhack-org/quack/backend/route"
	// Environment configuration and net/http server.
	"github.com/tamuhack-org/quack/backend/shared/server"
	// Reading in environmental variables.
	"github.com/tamuhack-org/quack/backend/shared/config"
	// MongoDB API implementation.
	"github.com/tamuhack-org/quack/backend/shared/database"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores, cuz its gon be lit
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load in environmental vars, which populate a config* struct.
	c, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal("There was an error reading in the config file\n", err)
	}

	// Connect to the database (this writes to a local variable in databse.go)
	err = database.CreateAndConnect(c.MongoUrl)
	if err != nil {
		log.Fatal("There was an error connecting with the database", err)
	}

	// Start the server, legggggggggooooooooo.
	server.Run(route.Load(), c)
}
