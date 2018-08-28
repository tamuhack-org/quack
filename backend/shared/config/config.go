package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	MongoUrl  string `json:"MongoUrl"`
	MongoName string `json:"MongoName"`
	Hostname  string `json:"Hostname"`
	SecretKey string `json:"SecretKey"`
	Port      int    `json:"Port"`
}

// Load the JSON config file.
func LoadFromEnv() (*Config, error) {
	// Set a default port if it hasn't been specified.
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	// User must specify a database url.
	mongoUrl := os.Getenv("MONGO_URL")
	if len(mongoUrl) <= 0 {
		return nil, errors.New("No database url specified!\n")
	}

	return &Config{
		Port:      port,
		MongoUrl:  mongoUrl,
	}, nil
}
