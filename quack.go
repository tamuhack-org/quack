package main

import (
	"log"
	// "os"
	"runtime"
  "fmt"

  // MVC-like routes.
	// "backend/shared/route"
  // MongoDB API implementation.
	"github.com/tamuhack-org/quack/backend/shared/database"
  // Our emailing logic (sendgrid?).
	// "backend/shared/email"
  // Environment configuration and net/http server.
	// "backend/shared/server"
  // Reading in environmental variables.
	"github.com/tamuhack-org/quack/backend/shared/config"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores, cuz its gon be lit
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load in environmental vars, which populate a config struct.
  c := config.LoadFromEnv()

	// Connect to database
  db, err := database.CreateAndConnect(c.MongoUrl)
	if err != nil {
    log.Fatal("There was an error connecting with the database", err)
	}


  fmt.Println(db)

	// Start the server, legggggggggooooooooo.
	// server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************
//
// config the settings variable
// var config = &configuration{}
//
// configuration contains the application settings
// type configuration struct {
//   Database  database.Info   `json:"Database"`
//   Server    server.Server   `json:"Server"`
//   Session   session.Session `json:"Session"`
// }
//
// ParseJSON unmarshals bytes to structs
// func (c *configuration) ParseJSON(b []byte) error {
//   return json.Unmarshal(b, &c)
// }
