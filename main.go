package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("MICRO_PORT")
	if port == "" {
		port = "80"
	}

	startServer(port)
}
