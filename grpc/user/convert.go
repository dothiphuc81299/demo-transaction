package grpcuser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-transaction/models"
	userpb "demo-transaction/proto/models/user"
)

func convertToUserBrief(data *userpb.UserBrief) models.UserBrief {
	var (
		userID, _ = primitive.ObjectIDFromHex(data.Id)
	)

	userBrief := models.UserBrief{
		ID:   userID,
		Name: data.Name,
		TotalTransaction:data.TotalTransaction,
		TotalCommission:data.TotalCommission,
	}
	return userBrief
}
