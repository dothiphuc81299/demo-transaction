package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// CompanyBrief ...
	CompanyBrief struct {
		ID               primitive.ObjectID `json:"_id"`
		Name             string             `json:"name"`
		CashbackPercent  float64            `json:"cashbackPercent"`
		TotalTransaction int64              `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue"`
	}
)
