package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	godotenv.Load(".env")

	value, err := godotenv.Read()

	if err != nil {
		log.Fatalf("%s", err)
	}

	return value[key]
}
