package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database_host     string
	Database_port     int
	Database_username string
	Database_name     string
	Database_password string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := Config{}

	if c.Database_host = os.Getenv("DATABASE_HOST"); c.Database_host == "" {
		c.Database_host = "localhost"
	}

	if c.Database_port, _ = strconv.Atoi(os.Getenv("DATABASE_PORT")); c.Database_port == 0 {
		c.Database_port = 27017
	}

	if c.Database_username = os.Getenv("DATABASE_USERNAME"); c.Database_username == "" {
		c.Database_username = ""
	}

	if c.Database_password = os.Getenv("DATABASE_PASSWORD"); c.Database_password == "" {
		c.Database_password = ""
	}

	if c.Database_name = os.Getenv("DATABASE_Name"); c.Database_name == "" {
		c.Database_name = ""
	}

	return c

}
