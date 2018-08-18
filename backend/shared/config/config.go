package config

import "os"

// Parser must implement ParseJSON
type Config struct {
	MongoUrl  string `json:"MongoUrl"`
	Hostname  string `json:"Hostname"`
	SecretKey string `json:"SecretKey"`
}

// Load the JSON config file
func LoadFromEnv() Config {
	return Config{
    MongoUrl: os.Getenv("MONGO_URL"),
    Hostname: os.Getenv("HOSTNAME"),
    SecretKey: os.Getenv("SECRET_KEY"),
  }
}
