package grpcnode

import (
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/dao"
	"demo-transaction/models"
	transactionpb "demo-transaction/proto/models/transaction"
	"demo-transaction/util"
)

func getTransactionDetailByUserID(userIDString string) ([]*transactionpb.TransactionDetail, error) {
	var (
		userID = util.HelperParseStringToObjectID(userIDString)
		filter = bson.M{"userID": userID}
	)

	// Get transactions
	transactionDetails, err := getTransactionDetailByFilter(filter)
	return transactionDetails, err
}

func getTransactionDetailByCompanyID(companyIDString string) ([]*transactionpb.TransactionDetail, error) {
	var (
		companyID = util.HelperParseStringToObjectID(companyIDString)
		filter    = bson.M{"companyID": companyID}
	)

	// Get transactions
	transactionDetails, err := getTransactionDetailByFilter(filter)
	return transactionDetails, err
}

func getTransactionDetailByFilter(filter bson.M) ([]*transactionpb.TransactionDetail, error) {
	var (
		transactionDetails = make([]*transactionpb.TransactionDetail, 0)
		wg                 sync.WaitGroup
	)

	// Find transactions
	transactions, err := dao.TransactionFindByFilter(filter)
	if err != nil {
		err = errors.New("Not Found Transaction by UserID")
		return nil, err
	}

	total := len(transactions)
	// Add process
	wg.Add(total)

	for index := range transactions {
		go func(index int) {
			defer wg.Done()

			// Convert to TransactionDetail
			transaction := convertToTransactionDetailGRPC(transactions[index])
			// Append
			transactionDetails = append(transactionDetails, transaction)
		}(index)
	}

	// Wait process
	wg.Wait()
	return transactionDetails, nil
}

func convertToTransactionDetailGRPC(transaction models.TransactionBSON) *transactionpb.TransactionDetail {
	var (
		idString        = transaction.ID.Hex()
		companyIDString = transaction.CompanyID.Hex()
		branchIDString  = transaction.BranchID.Hex()
		userIDString    = transaction.UserID.Hex()
		createdAt       = util.HelperConvertTimeToTimestampProto(transaction.CreatedAt)
	)

	// TransactionDetail
	result := &transactionpb.TransactionDetail{
		Id:                     idString,
		CompanyID:              companyIDString,
		BranchID:               branchIDString,
		UserID:                 userIDString,
		Amount:                 transaction.Amount,
		Commission:             transaction.Commission,
		CompanyCashbackPercent: transaction.CompanyCashbackPercent,
		CreatedAt:              createdAt,
	}
	return result
}
