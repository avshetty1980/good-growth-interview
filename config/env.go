package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	Username   string
	Password   string
	DBURI      string
	DBName     string
	Collection string
}

var Envs = InitConfig()

func InitConfig() Config {
	return Config{
		Port:       getEnv("PORT", ":8080"),
		Username:   getEnv("MONGO_DB_USERNAME", ":8080"),
		Password:   getEnv("MONGO_DB_PASSWORD", ":8080"),
		DBURI:      fmt.Sprintf("mongodb://%s:%s@db:%s/", getEnv("DB_USER", "mongoadmin"), getEnv("DB_PASSWORD", "secret"), getEnv("DB_PORT", "27017")),
		DBName:     getEnv("DB_NAME", "goodgrowth"),
		Collection: getEnv("Collection", "messages"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
