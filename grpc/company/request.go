package grpccompany

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"demo-transaction/models"
// 	companypb "demo-transaction/proto/models/company"
// )

// // GetCompanyBriefByID ...
// func GetCompanyBriefByID(companyID string) (companyBrief models.CompanyBrief, err error) {
// 	clientConn, client := CreateClient()
// 	defer clientConn.Close()

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	// Call
// 	result, err := client.GetCompanyBriefByID(ctx, &userpb.GetCompanyBriefByIDRequest{Id: companyID})
// 	if err != nil {
// 		err = errors.New("Khong the get company brief by id")
// 		return
// 	}

// 	// Convert to user brief
// 	companyBrief := convertToUserBrief(result.CompanyBrief)

// 	return companyBrief, nil
// }
