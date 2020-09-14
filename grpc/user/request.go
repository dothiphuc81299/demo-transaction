package grpcuser

import (
	"log"
	"context"
	"time"

	"demo-transaction/models"
	userpb "demo-transaction/proto/models/user"
)

// GetUserBriefByID ...
func GetUserBriefByID(userID string) (userBrief models.UserBrief, err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call GetUserBriefByID
	result, err := client.GetUserBriefByID(ctx, &userpb.GetUserBriefByIDRequest{UserID: userID})
	if err != nil {
		log.Printf("Call grpc get user by Id error %v\n", err)
		return 
	}

	// Convert to user brief
	userBrief = convertToUserBrief(result.UserBrief)
	return
}

// UpdateUserStatsByID ...
func UpdateUserStatsByID(userID string, totalTransaction int64, totalCommission float64) (err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call UpdateUserStatsByID
	_, err = client.UpdateUserStatsByID(ctx, &userpb.UpdateUserStatsByIDRequest{
		Id:               userID,
		TotalTransaction: totalTransaction,
		TotalCommission:  totalCommission,
	})
	if err != nil {
		log.Printf("Call grpc update userStats error %v\n", err)
		return
	}
	return
}
