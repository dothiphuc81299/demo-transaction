package models

import(
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (

	// TransactionCreatePayload ...
	TransactionCreatePayload struct {
		CompanyID string  `json:"companyID"`
		BranchID  string  `json:"branchID"`
		UserID    string  `json:"userID"`
		Amount    float64 `json:"amount"`
	}
)

// Validate TransactionCreatePayload
func (payload TransactionCreatePayload) Validate() error {
	err := validation.Errors{
		"companyID": validation.Validate(payload.CompanyID, validation.Required, is.MongoID),
		"branchID": validation.Validate(payload.BranchID, validation.Required, is.MongoID),
		"userID": validation.Validate(payload.UserID, validation.Required, is.MongoID),
		"amount": validation.Validate(payload.Amount, validation.Required,is.Digit),
	}.Filter()
	return err
}