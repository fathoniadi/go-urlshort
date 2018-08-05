package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	database_host     string
	database_port     int
	database_username string
	database_password string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := Config{}

	if c.database_host = os.Getenv("DATABASE_HOST"); c.database_host == "" {
		c.database_host = "localhost"
	}

	if c.database_port, _ = strconv.Atoi(os.Getenv("DATABASE_PORT")); c.database_port == 0 {
		c.database_port = 27017
	}

	if c.database_username = os.Getenv("DATABASE_USERNAME"); c.database_username == "" {
		c.database_username = ""
	}

	if c.database_password = os.Getenv("DATABASE_PASSWORD"); c.database_password == "" {
		c.database_password = ""
	}

	return c

}
