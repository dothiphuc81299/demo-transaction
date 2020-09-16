package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// TransactionCreatePayload ...
	TransactionCreatePayload struct {
		Company      string             `json:"company"`
		CompanyID    primitive.ObjectID `json:"companyID"`
		CompanyBrief CompanyBrief       `json:"companyBrief"`
		Branch       string             `json:"branch"`
		BranchID     primitive.ObjectID `json:"branchID"`
		BranchBrief  BranchBrief        `json:"branchBrief"`
		User         string             `json:"user"`
		UserID       primitive.ObjectID `json:"userID"`
		UserBrief    UserBrief          `json:"userBrief"`
		Amount       float64            `json:"amount"`
	}
)

// Validate TransactionCreatePayload
func (payload TransactionCreatePayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Company,
			validation.Required.Error("company is required"),
			is.MongoID.Error("company is MongoID"),
		),
		validation.Field(
			&payload.Branch,
			validation.Required.Error("branch is required"),
			is.MongoID.Error("branch is MongoID"),
		),
		validation.Field(
			&payload.User,
			validation.Required.Error("user is required"),
			is.MongoID.Error("user is MongoID"),
		),
		validation.Field(
			&payload.Amount,
			validation.Required.Error("amount is required"),
		),
	)
}

// ConvertToBson ...
func (payload TransactionCreatePayload) ConvertToBson() TransactionBSON {
	result := TransactionBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: payload.CompanyID,
		BranchID:  payload.BranchID,
		UserID:    payload.UserID,
		Amount:    payload.Amount,
		CreatedAt: time.Now(),
	}
	return result
}
