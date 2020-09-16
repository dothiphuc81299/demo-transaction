package grpcuser

import (
	"demo-transaction/models"
	userpb "demo-transaction/proto/models/user"
	"demo-transaction/utils"
)

func convertToUserBrief(data *userpb.UserBrief) models.UserBrief {
	var (
		userID = utils.HelperParseStringToObjectID(data.Id)
	)

	// UserBrief struct
	userBrief := models.UserBrief{
		ID:               userID,
		Name:             data.Name,
		TotalTransaction: data.TotalTransaction,
		TotalCommission:  data.TotalCommission,
	}
	return userBrief
}
