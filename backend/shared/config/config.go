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
	mongoUrl, err := strconv.Atoi(os.Getenv("MONGO_URL"))
	if err != nil {
		return nil, errors.New("No database url specified!")
	}

	return &Config{
		Port:      port,
		MongoUrl:  os.Getenv("MONGO_URL"),
		MongoName: os.Getenv("MONGO_NAME"),
		Hostname:  os.Getenv("HOSTNAME"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}, nil
}
