package config

import (
	"errors"
	"log"
	"os"
	"strconv"
)

// This is a persistent session that's accessed by our handlers.
var (
	GlobalConfig *Config
)

type Config struct {
	MongoUrl        string `json:"MongoUrl"`
	Port            int    `json:"Port"`
	EventbriteUrl   string `json:"EventbriteUrl"`
	EventbriteToken string `json:"EventbriteToken"`
}

// Load the JSON config file.
func LoadFromEnv() error {
	// Set a default port if it hasn't been specified.
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	// User must specify a database url.
	mongoUrl := os.Getenv("MONGO_URL")
	if len(mongoUrl) <= 0 {
		return errors.New("No database url specified!\n")
	}

	// Eventbrite url and token is optional, but it'll be logged.
	eventbriteUrl := os.Getenv("EVENTBRITE_URL")
	eventbriteToken := os.Getenv("EVENTBRITE_TOKEN")
	if len(eventbriteUrl) <= 0 || len(eventbriteToken) <= 0 {
		log.Println("Empty eventbrite metadata specified")
	}

	log.Println("----------------------------------------------------------")
	log.Println("Starting server on port :" + strconv.Itoa(port) + " with url " + mongoUrl)
	log.Println("Eventbrite Url is " + eventbriteUrl + " and Eventbrite Token is " + eventbriteToken)
	log.Println("----------------------------------------------------------")

	// Assign.
	GlobalConfig =
		&Config{
			Port:            port,
			MongoUrl:        mongoUrl,
			EventbriteUrl:   eventbriteUrl,
			EventbriteToken: eventbriteToken,
		}
	return nil
}
