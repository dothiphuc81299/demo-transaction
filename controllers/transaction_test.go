package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/modules/database"
	"demo-transaction/util"
	"demo-transaction/models"
	"demo-transaction/modules/redis"
)

type TransactionSuite struct {
	suite.Suite
}

func (s TransactionSuite) SetupSuite() {
	util.HelperConnect()
	redis.Connect()
	util.HelperCompanyCreateFake()
	util.HelperLoyaltyProgramFake()
	util.HelperBranchCreateFake()
	util.HelperUserCreateFake()
	util.HelperCompanyAnalyticFake()
}

func (s TransactionSuite) TearDownSuite() {
	removeOldDataTransaction()
}
func removeOldDataTransaction() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
	database.LoyaltyProgramCol().DeleteMany(context.Background(), bson.M{})
	database.BranchCol().DeleteMany(context.Background(), bson.M{})
	database.UserCol().DeleteMany(context.Background(), bson.M{})
	database.TransactionCol().DeleteMany(context.Background(), bson.M{})
	database.TransactionAnalyticCol().DeleteMany(context.Background(), bson.M{})
	database.LoyaltyProgramUserStatusCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyAnalyticCol().DeleteMany(context.Background(), bson.M{})
}

func (s *TransactionSuite) TestTransactionCreatePrepaidSuccess() {
	var (
		transaction = models.TransactionCreatePayload{
			CompanyID: util.CompanyString,
			BranchID:  util.BranchString,
			UserID:    util.UserString,
			Amount:    100000,
		}
		response util.Response
	)

	// Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", transaction)
	c.Set("company", util.Company)
	c.Set("branch", util.Branch)
	c.Set("user", util.User)

	// Call TransactionCreate
	TransactionCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	// Get totalTransaction
	transactionAnalytic := util.HelperTransactionAnalyticFindByID()
	totalTransaction := transactionAnalytic.TotalTransaction

	//Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
	assert.Equal(s.T(), 1, totalTransaction)
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
