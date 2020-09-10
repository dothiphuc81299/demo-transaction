package routes

import (
	"demo-transaction/models"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	grpcuser "demo-transaction/grpc/user"
	//grpccompany "demo-transaction/grpc/company"
	"demo-transaction/util"
)

// func companyCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var (
// 			companyID     = c.Get("companyID").(primitive.ObjectID)
// 			companyString = companyID.Hex()
// 		)

// 		user, err := grpcuser.GetCompanyBriefByID(companyString)
// 		if err != nil {
// 			return util.Response404(c, nil, "Khong tim thay user")
// 		}

// 		c.Set("company", user)
// 		return next(c)
// 	}
// }

// func branchCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var (
// 			userID     = c.Get("userID").(primitive.ObjectID)
// 			userString = userID.Hex()
// 		)

// 		user, err := grpcuser.GetUserBriefByID(userString)
// 		if err != nil {
// 			return util.Response404(c, nil, "Khong tim thay user")
// 		}

// 		c.Set("branch", user)
// 		return next(c)
// 	}
// }
// Validate company object id

func userCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body    = c.Get("body").(models.TransactionCreatePayload)
			userID = body.UserID
		)

		user, err := grpcuser.GetUserBriefByID(userID)
		if err != nil {
			return util.Response404(c, nil, "Not Found User by ID")
		}

		companyID     := c.Get("companyID").(primitive.ObjectID)
		company:=models.CompanyBrief{
			ID             :  companyID,
			Name           : "hoang",
			CashbackPercent  : 10,
			TotalTransaction : 0,
			TotalRevenue   :0,
		}

		branchID     := c.Get("branchID").(primitive.ObjectID)
		branch:=models.BranchBrief{
			ID             :  branchID,
			Name           : "hoang",
			TotalTransaction  : 0,
			TotalRevenue : 0,
		}


		c.Set("company", company)
		c.Set("branch", branch)
		c.Set("user", user)
		return next(c)
	}

}
