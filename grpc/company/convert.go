package grpccompany

import (
	"demo-transaction/models"
	companypb "demo-transaction/proto/models/company"
	"demo-transaction/utils"
)

func convertToCompanyBrief(data *companypb.CompanyBrief) models.CompanyBrief {
	var (
		companyID = utils.HelperParseStringToObjectID(data.Id)
	)

	companyBrief := models.CompanyBrief{
		ID:               companyID,
		Name:             data.Name,
		CashbackPercent:  data.CashbackPercent,
		TotalTransaction: data.TotalTransaction,
		TotalRevenue:     data.TotalRevenue,
	}
	return companyBrief
}

func convertToBranchBrief(data *companypb.BranchBrief) models.BranchBrief {
	var (
		branchID = utils.HelperParseStringToObjectID(data.Id)
	)

	branchBrief := models.BranchBrief{
		ID:               branchID,
		Name:             data.Name,
		TotalTransaction: data.TotalTransaction,
		TotalRevenue:     data.TotalRevenue,
	}
	return branchBrief
}
