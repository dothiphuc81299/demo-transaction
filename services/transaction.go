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
	)

	// Check User Request
	err = transactionCheckUserRequest(userString)
	if err != nil {
		return
	}
	redis.Set(config.RedisKeyUser, userString)

	// Calculate commission
	commission := calculateTransactionCommison(companyBrief.CashbackPercent, body.Amount)

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
	err = userUpdateAfterCreateTransaction(transaction,userBrief,commission)
	if err != nil {
		return
	}
	return
}

// TransactionFindByUserID ...
// func TransactionFindByUserID(user models.UserBrief) ([]models.TransactionDetail, error) {
// 	var (
// 		result = make([]models.TransactionDetail, 0)
// 		wg     sync.WaitGroup
// 	)

// 	// Find
// 	transactions, err := dao.TransactionFindByUserID(user.ID)
// 	total := len(transactions)

// 	// Return if not found
// 	if total == 0 {
// 		return result, err
// 	}

// 	// Add process
// 	wg.Add(total)

// 	for index := range transactions {
// 		go func(index int) {
// 			defer wg.Done()

// 			// Convert to TransactionDetail
// 			transaction := convertToTransactionDetail(transactions[index], user)

// 			// Append
// 			result = append(result, transaction)
// 		}(index)
// 	}

// 	// Wait process
// 	wg.Wait()

// 	// Sort again
// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].CreatedAt.After(result[j].CreatedAt)
// 	})

// 	return result, err
// }
