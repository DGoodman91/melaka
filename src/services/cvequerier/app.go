package main

import (
	"context"
	"fmt"
	"log"
)

func main() {

	// Create a DB instance and connect to our database
	// Note that we're not doing this inside an init function because the init function is executed on test runs

	db := &MongoDB{
		Configuration: DBConnConfig{
			Url:            readFromENV("MONGO_URL", "mongodb://localhost:27017"),
			Username:       readFromENV("MONGO_ROOT_USERNAME", "dev"),
			Password:       readFromENV("MONGO_ROOT_PASSWORD", "dev"),
			Database:       readFromENV("MONGO_DATABASE", "melakaDB"),
			CveCollection:  readFromENV("MONGO_CVE_COLLECTION", "cves"),
			MetaCollection: readFromENV("MONGO_META_COLLECTION", "meta"),
		},
	}
	defer db.Connection.Disconnect(context.Background())

	err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to DB instance with error: %s", err)
	}

	// Set up our server with it's routes & middleware
	server := buildServer(db)

	// Run our server!
	port := readFromENV("LISTEN_PORT", "8080")
	server.Run(fmt.Sprintf(":%s", port))
	fmt.Printf("Listening on port %s\n", port)

}
