package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"demo-transaction/config"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

// Connect ...
func Connect() {
	envVars := config.GetEnv()

	// Set Client
	client, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot create client with uri:", envVars.Database.URI)
	}

	// Set deadline
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", envVars.Database.URI)
	}

	// Set database
	db = client.Database(envVars.Database.Name)
	fmt.Println("Database Connected to", envVars.Database.Name)
}

// GetClient ...
func GetClient() *mongo.Client {
	return client
}

// SetDB ...
func SetDB(dbValue *mongo.Database) {
	db = dbValue
}

// GetDB ...
func GetDB() *mongo.Database {
	return db
}
