package services

import (
	"demo-transaction/config"
	"demo-transaction/dao"
	"demo-transaction/models"
	"demo-transaction/modules/redis"
)

// TransactionCreate ...
func TransactionCreate(body models.TransactionCreatePayload, companyBrief models.CompanyBrief, branchBrief models.BranchBrief, userBrief models.UserBrief) (transaction models.TransactionBSON, err error) {
	var (
		userString = body.UserID
		companyID  = companyBrief.ID
		branchID   = branchBrief.ID
		userID     = userBrief.ID
		amount     = body.Amount
	)

	// Check User Request
	err = transactionCheckUserRequest(userString)
	if err != nil {
		return
	}
	redis.Set(config.RedisKeyUser, userString)

	// Calculate commission
	commission := calculateTransactionCommison(companyBrief.CashbackPercent, amount)

	// Convert Transaction
	transaction = transactionCreatePayloadToBSON(body, companyID, branchID, userID)

	// Add information for Transaction
	transaction = transactionAddInformation(transaction, commission, companyBrief.CashbackPercent)

	// Create Transaction
	_, err = dao.TransactionCreate(transaction)
	if err != nil {
		return
	}

	// Update Company
	err = companyUpdateAfterCreateTransaction(companyBrief, amount)
	if err != nil {
		return
	}

	// Update Branch
	err = branchUpdateAfterCreateTransaction(branchBrief, amount)
	if err != nil {
		return
	}

	// Update User
	err = userUpdateAfterCreateTransaction(userBrief, commission)
	if err != nil {
		return
	}
	return
}
