package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// TransactionBSON ...
	TransactionBSON struct {
		ID                     primitive.ObjectID `bson:"_id"`
		CompanyID              primitive.ObjectID `bson:"companyID"`
		BranchID               primitive.ObjectID `bson:"branchID"`
		UserID                 primitive.ObjectID `bson:"userID"`
		Amount                 float64            `bson:"amount"`
		Commission             float64            `bson:"commission"`
		CompanyCashbackPercent float64            `bson:"companyCashbackPercent"`
		CreatedAt              time.Time          `bson:"createdAt"`
	}

	//TransactionDetail ...
	TransactionDetail struct {
		ID                     primitive.ObjectID `json:"_id"`
		Company                string             `json:"company"`
		Branch                 string             `json:"branch"`
		User                   string             `json:"user"`
		Amount                 float64            `json:"amount"`
		Commission             float64            `json:"commission"`
		CompanyCashbackPercent float64            `json:"companyCashbackPercent"`
		CreatedAt              time.Time          `json:"createdAt"`
	}
)
