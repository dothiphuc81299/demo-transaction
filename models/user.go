package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// UserBrief ...
	UserBrief struct {
		ID               primitive.ObjectID `json:"_id"`
		Name             string             `json:"name"`
		TotalTransaction int64              `json:"totalTransaction"`
		TotalCommission  float64            `json:"totalCommission"`
	}
)
