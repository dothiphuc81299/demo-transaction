package services

import (
	"demo-transaction/util"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-transaction/config"
	grpccompany "demo-transaction/grpc/company"
	grpcuser "demo-transaction/grpc/user"
	"demo-transaction/models"
	"demo-transaction/modules/redis"
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
	commission := (CompanyCashbackPercent / 100) * amount
	return commission
}

func transactionAddInformation(transaction models.TransactionBSON, commission, companyCashbackPercent float64) models.TransactionBSON {
	transaction.Commission = commission
	transaction.CompanyCashbackPercent = companyCashbackPercent
	transaction.ID = primitive.NewObjectID()
	transaction.CreatedAt = time.Now()
	return transaction
}

func companyUpdateAfterCreateTransaction(companyBrief models.CompanyBrief, amount float64) (err error) {
	var (
		companyID        = companyBrief.ID
		companyIDString  = util.HelperParseObjectIDToString(companyID)
		totalTransaction = companyBrief.TotalTransaction
		totalRevenue     = companyBrief.TotalRevenue
	)

	// Set userStats
	totalTransaction++
	totalRevenue += amount
	err = grpccompany.UpdateCompanyStatsByID(companyIDString, totalTransaction, totalRevenue)
	return
}

func branchUpdateAfterCreateTransaction(branchBrief models.BranchBrief, amount float64) (err error) {
	var (
		branchID         = branchBrief.ID
		branchIDString   = util.HelperParseObjectIDToString(branchID)
		totalTransaction = branchBrief.TotalTransaction
		totalRevenue     = branchBrief.TotalRevenue
	)

	// Set userStats
	totalTransaction++
	totalRevenue += amount
	err = grpccompany.UpdateBranchStatsByID(branchIDString, totalTransaction, totalRevenue)
	return
}

func userUpdateAfterCreateTransaction(userBrief models.UserBrief, commission float64) (err error) {
	var (
		userID           = userBrief.ID
		userIDString     = util.HelperParseObjectIDToString(userID)
		totalTransaction = userBrief.TotalTransaction
		totalCommission  = userBrief.TotalCommission
	)

	// Set userStats
	totalTransaction++
	totalCommission += commission
	err = grpcuser.UpdateUserStatsByID(userIDString, totalTransaction, totalCommission)
	return
}
