package database

import (
	"context"
	"fmt"
	"log"

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

	// Connect Client
	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println("err", err)
		log.Fatal("Cannot connect to database ", envVars.Database.URI)
	}

	// Set client
	client = cl

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
