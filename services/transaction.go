package services

import (
	"demo-transaction/config"
	"demo-transaction/dao"
	"demo-transaction/models"
	"demo-transaction/modules/redis"
)

// TransactionCreate ...
func TransactionCreate(payload models.TransactionCreatePayload) (transaction models.TransactionBSON, err error) {
	var (
		userIDString = payload.User
		companyBrief = payload.CompanyBrief
		branchBrief  = payload.BranchBrief
		userBrief    = payload.UserBrief
		amount       = payload.Amount
	)

	// Check user request
	err = transactionCheckUserRequest(userIDString)
	if err != nil {
		return
	}
	redis.Set(config.RedisKeyUser, userIDString)

	// Calculate commission
	commission := calculateTransactionCommison(companyBrief.CashbackPercent, amount)

	// Convert to TransactionBSON
	transaction = payload.ConvertToBson()

	// Add information for Transaction
	transaction = transactionAddInformation(transaction, commission, companyBrief.CashbackPercent)

	// Create transaction
	_, err = dao.TransactionCreate(transaction)
	if err != nil {
		return
	}

	// Update company
	err = companyUpdateAfterCreateTransaction(companyBrief, amount)
	if err != nil {
		return
	}

	// Update branch
	err = branchUpdateAfterCreateTransaction(branchBrief, amount)
	if err != nil {
		return
	}

	// Update user
	err = userUpdateAfterCreateTransaction(userBrief, commission)
	if err != nil {
		return
	}

	return
}
