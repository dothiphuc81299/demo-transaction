package grpcnode

import (
	"demo-transaction/models"
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/dao"
	transactionpb "demo-transaction/proto/models/transaction"
	"demo-transaction/util"
)

func getTransactionDetailByUserID(userIDString string) (*transactionpb.GetTransactionDetailByUserIDResponse, error) {
	var (
		userID             = util.HelperParseStringToObjectID(userIDString)
		filter             = bson.M{"_id": userID}
		transactionDetails = make([]*transactionpb.TransactionDetail, 0)
		wg                 sync.WaitGroup
	)

	// Find Transactions
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
			transaction := convertToTransactionDetail(transactions[index])

			// Append
			transactionDetails = append(transactionDetails, transaction)
		}(index)
	}

	// Wait process
	wg.Wait()

	// Success
	result := &transactionpb.GetTransactionDetailByUserIDResponse{
		TransactionDetail: transactionDetails,
	}

	return result, nil
}

func convertToTransactionDetail(transaction models.TransactionBSON) *transactionpb.TransactionDetail {
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
