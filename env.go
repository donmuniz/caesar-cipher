package main

import (
	"log"
	"os"
)

func getenv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatal(key + " must be set")
	}
	return value
}
