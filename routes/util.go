package routes

import (
	"github.com/labstack/echo/v4"

	grpccompany "demo-transaction/grpc/company"
	grpcuser "demo-transaction/grpc/user"
	"demo-transaction/models"
	"demo-transaction/util"
)

func companyCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body      = c.Get("body").(models.TransactionCreatePayload)
			companyID = body.CompanyID
		)

		// Call grpc api
		companyBrief, err := grpccompany.GetCompanyBriefByID(companyID)
		if err != nil {
			return util.Response404(c, nil, "Not found company by ID")
		}

		// Set value
		c.Set("companyBrief", companyBrief)
		return next(c)
	}
}

func branchCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body     = c.Get("body").(models.TransactionCreatePayload)
			branchID = body.BranchID
		)

		// Call grpc api
		branchBrief, err := grpccompany.GetBranchBriefByID(branchID)
		if err != nil {
			return util.Response404(c, nil, "Not found banch by ID")
		}

		// Set value
		c.Set("branchBrief", branchBrief)
		return next(c)
	}
}

func userCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body   = c.Get("body").(models.TransactionCreatePayload)
			userID = body.UserID
		)

		// Call grpc api
		userBrief, err := grpcuser.GetUserBriefByID(userID)
		if err != nil {
			return util.Response404(c, nil, "Not found user by ID")
		}

		// Set value
		c.Set("userBrief", userBrief)
		return next(c)
	}
}
