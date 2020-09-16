package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/models"
	"demo-transaction/services"
	"demo-transaction/utils"
)

// TransactionCreate ...
func TransactionCreate(c echo.Context) error {
	var (
		payload = c.Get("payload").(models.TransactionCreatePayload)
	)

	// Process data
	rawData, err := services.TransactionCreate(payload)

	// If err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}
