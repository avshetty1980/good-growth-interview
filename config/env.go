package config

import (
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
		Username:   getEnv("MONGO_DB_USERNAME", "mongoadmin"),
		Password:   getEnv("MONGO_DB_PASSWORD", "secret"),
		DBURI:      getEnv("MONGO_DB_URI", "mongodb://mongoadmin:secret@db:27017/"),
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
