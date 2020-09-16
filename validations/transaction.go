package validations

import (
	"github.com/labstack/echo/v4"

	grpccompany "demo-transaction/grpc/company"
	grpcuser "demo-transaction/grpc/user"
	"demo-transaction/models"
	"demo-transaction/utils"
)

// TransactionCreate ...
func TransactionCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload models.TransactionCreatePayload
		)

		// ValidateStruct
		c.Bind(&payload)
		err := payload.Validate()

		//if err
		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// Check exit company by call gRPC
		companyBrief, err := grpccompany.GetCompanyBriefByID(payload.Company)
		if companyBrief.ID.IsZero() {
			return utils.Response404(c, nil, err.Error())
		}

		// Check exit branch by call gRPC
		branchBrief, err := grpccompany.GetBranchBriefByID(payload.Branch)
		if branchBrief.ID.IsZero() {
			return utils.Response404(c, nil, err.Error())
		}
		// Check exit user by call gRPC
		userBrief, err := grpcuser.GetUserBriefByID(payload.User)
		if userBrief.ID.IsZero() {
			return utils.Response404(c, nil, err.Error())
		}

		// Add information for payload
		payload.CompanyID = companyBrief.ID
		payload.CompanyBrief = companyBrief
		payload.BranchID = branchBrief.ID
		payload.BranchBrief = branchBrief
		payload.UserID = userBrief.ID
		payload.UserBrief = userBrief

		// Success
		c.Set("payload", payload)
		return next(c)
	}
}
