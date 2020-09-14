package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// BranchBrief ...
	BranchBrief struct {
		ID               primitive.ObjectID `json:"_id"`
		Name             string             `json:"name"`
		TotalTransaction int64              `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue"`
	}
)
