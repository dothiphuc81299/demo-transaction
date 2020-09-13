package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/models"
	"demo-transaction/modules/database"
)

// TransactionCreate ...
func TransactionCreate(doc models.TransactionBSON) (models.TransactionBSON, error) {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)

	// Insert
	_, err := collection.InsertOne(ctx, doc)
	return doc, err
}

// TransactionFindByFilter ...
func TransactionFindByFilter(filter bson.M) ([]models.TransactionBSON, error) {
	var (
		transactionCol = database.TransactionCol()
		ctx            = context.Background()
		result         = make([]models.TransactionBSON, 0)
	)

	// Find
	cursor, err := transactionCol.Find(ctx, filter)

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)
	return result, err
}
