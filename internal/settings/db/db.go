package db

import "os"

var (
	// Port is the port the server will listen on
	Port string = ":" + os.Getenv("PORT")
	// DatabaseURL is the URL to the database
	DatabaseURL string = os.Getenv("DATABASE_URL")
)
