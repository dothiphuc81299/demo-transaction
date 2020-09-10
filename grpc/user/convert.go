package grpcuser

import (
	"demo-transaction/util"

	"demo-transaction/models"
	userpb "demo-transaction/proto/models/user"
)

func convertToUserBrief(data *userpb.UserBrief) models.UserBrief {
	var (
		userID = util.HelperParseStringToObjectID(data.Id)
	)

	userBrief := models.UserBrief{
		ID:               userID,
		Name:             data.Name,
		TotalTransaction: data.TotalTransaction,
		TotalCommission:  data.TotalCommission,
	}
	return userBrief
}
