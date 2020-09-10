package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection name
const (
	transactions = "transactions"
)

// TransactionCol ...
func TransactionCol() *mongo.Collection {
	return db.Collection(transactions)
}
