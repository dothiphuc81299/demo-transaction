package grpccompany

import (
	"context"
	"log"
	"time"

	"demo-transaction/models"
	companypb "demo-transaction/proto/models/company"
)

// GetCompanyBriefByID ...
func GetCompanyBriefByID(companyID string) (companyBrief models.CompanyBrief, err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call GetCompanyBriefByID
	result, err := client.GetCompanyBriefByID(ctx, &companypb.GetCompanyBriefByIDRequest{CompanyID: companyID})
	if err != nil {
		log.Printf("Call grpc get Company by Id error %v\n", err)
		return
	}

	// Convert to companyBrief
	companyBrief = convertToCompanyBrief(result.CompanyBrief)
	return
}

// GetBranchBriefByID ...
func GetBranchBriefByID(branchID string) (branchBrief models.BranchBrief, err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call GetBranchBriefByID
	result, err := client.GetBranchBriefByID(ctx, &companypb.GetBranchBriefByIDRequest{BranchID: branchID})
	if err != nil {
		log.Printf("Call grpc get Branch by Id error %v\n", err)
		return
	}

	// Convert to Branch brief
	branchBrief = convertToBranchBrief(result.BranchBrief)
	return
}

// UpdateCompanyStatsByID ...
func UpdateCompanyStatsByID(companyID string, totalTransaction int64, totalRevenue float64) (err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call UpdateCompanyStatsByID
	_, err = client.UpdateCompanyStatsByID(ctx, &companypb.UpdateCompanyStatsByIDRequest{
		Id:               companyID,
		TotalTransaction: totalTransaction,
		TotalRevenue:     totalRevenue,
	})
	if err != nil {
		log.Printf("Call grpc update CompanyStats error %v\n", err)
		return
	}
	return
}

// UpdateBranchStatsByID ...
func UpdateBranchStatsByID(branchID string, totalTransaction int64, totalRevenue float64) (err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call UpdateBranchStatsByID
	_, err = client.UpdateBranchStatsByID(ctx, &companypb.UpdateBranchStatsByIDRequest{
		Id:               branchID,
		TotalTransaction: totalTransaction,
		TotalRevenue:     totalRevenue,
	})
	if err != nil {
		log.Printf("Call grpc update BranchStats error %v\n", err)
		return
	}
	return
}
