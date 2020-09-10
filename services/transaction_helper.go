package services

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-transaction/config"
	"demo-transaction/models"
	"demo-transaction/modules/redis"
	grpcuser "demo-transaction/grpc/user"
)

func transactionCreatePayloadToBSON(body models.TransactionCreatePayload, companyID, branchID, userID primitive.ObjectID) models.TransactionBSON {

	result := models.TransactionBSON{
		CompanyID: companyID,
		BranchID:  branchID,
		UserID:    userID,
		Amount:    body.Amount,
	}

	return result
}

func transactionCheckUserRequest(userString string) (err error) {
	userReq := redis.Get(config.RedisKeyUser)
	if userReq == userString {
		err = errors.New("User Dang Thuc hien giao dich")
		return
	}
	return
}

func calculateTransactionCommison(CompanyCashbackPercent, amount float64) float64 {
	var (
		commission float64
	)
	commission = (CompanyCashbackPercent / 100) * amount

	return commission
}

func transactionAddInformation(transaction models.TransactionBSON, commission, companyCashbackPercent float64) models.TransactionBSON {
	transaction.Commission = commission
	transaction.CompanyCashbackPercent = companyCashbackPercent
	transaction.ID = primitive.NewObjectID()
	transaction.CreatedAt = time.Now()
	return transaction
}

// func companyAndBranchUpdateAfterCreateTransaction(transaction models.TransactionBSON) (err error) {
// 	var (
// 		companyID     = transaction.CompanyID
// 		userID     = transaction.UserID
// 	)
// }

func userUpdateAfterCreateTransaction(transaction models.TransactionBSON, userBrief models.UserBrief,commission float64) (err error) {
	var (
		userID     = transaction.UserID
		userIDString = userID.Hex()
		totalTransaction = userBrief.TotalTransaction
		totalCommission = userBrief.TotalCommission 
	)

	// Set userStats
	totalTransaction++
	totalCommission += commission

	err = grpcuser.UpdateUserStatsByID(userIDString,totalTransaction,totalCommission)
	return 
}


// func convertToTransactionDetail(transaction models.TransactionBSON, user models.UserBrief) models.TransactionDetail {
// 	var (
// 		company, _  = dao.CompanyFindByID(transaction.CompanyID)
// 		branch, _   = dao.BranchFindByID(transaction.BranchID)
// 		companyName = company.Name
// 		branchName  = branch.Name
// 		userName    = user.Name
// 	)

// 	// TransactionDetail
// 	result := models.TransactionDetail{
// 		ID:                       transaction.ID,
// 		Company:                  companyName,
// 		Branch:                   branchName,
// 		User:                     userName,
// 		Amount:                   transaction.Amount,
// 		Commission:               transaction.Commission,
// 		CompanyCashbackPercent:   transaction.CompanyCashbackPercent,
// 		MilestoneCashbackPercent: transaction.MilestoneCashbackPercent,
// 		PaidType:                 transaction.PaidType,
// 		CreatedAt:                transaction.CreatedAt,
// 	}

// 	return result
// }
