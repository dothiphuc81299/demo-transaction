package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/models"
	"demo-transaction/services"
	"demo-transaction/util"
)

// TransactionCreate ...
func TransactionCreate(c echo.Context) error {
	var (
		body    = c.Get("body").(models.TransactionCreatePayload)
		company = c.Get("company").(models.CompanyBrief)
		branch  = c.Get("branch").(models.BranchBrief)
		user    = c.Get("user").(models.UserBrief)
	)

	// Process data
	rawData, err := services.TransactionCreate(body, company, branch, user)

	//if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")

}