package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"

	"demo-transaction/apptest"
	"demo-transaction/models"
	"demo-transaction/modules/database"
	"demo-transaction/utils"
)

// Test Create Transaction
type TransactionCreateTestSuite struct {
	suite.Suite
	e    *echo.Echo
	data models.TransactionCreatePayload
}

func (suite *TransactionCreateTestSuite) SetupSuite() {
	// Init server
	suite.e = apptest.InitServer()

	// Clear Data
	removeOldDataTransaction()

	// Setup payload data
	suite.data = setupDataTransaction()
}

func (suite *TransactionCreateTestSuite) TearDownSuite() {
	removeOldDataTransaction()
}

func (suite *TransactionCreateTestSuite) TestTransactionCreateSuccess() {
	var (
		payload      = suite.data
		response     utils.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/hoang/Documents/Company/demo-transaction/schemas/transaction_create.json")
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Create JSONLoader from go struct
	documentLoader := gojsonschema.NewGoLoader(response)

	// Validate json response
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	// Test
	assert.Equal(suite.T(), true, result.Valid())
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

func (suite *TransactionCreateTestSuite) TestTransactionCreateFailureWithInvalidCompanyID() {
	var (
		payload = models.TransactionCreatePayload{
			Company: "1",
			Branch:  suite.data.Branch,
			User:    suite.data.User,
			Amount:  suite.data.Amount,
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	log.Println("respone:", response)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *TransactionCreateTestSuite) TestTransactionCreateFailureWithInvalidBranchID() {
	var (
		payload = models.TransactionCreatePayload{
			Company: suite.data.Company,
			Branch:  "1",
			User:    suite.data.User,
			Amount:  suite.data.Amount,
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *TransactionCreateTestSuite) TestTransactionCreateFailureWithInvalidUserID() {
	var (
		payload = models.TransactionCreatePayload{
			Company: suite.data.Company,
			Branch:  suite.data.Branch,
			User:    "1",
			Amount:  suite.data.Amount,
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *TransactionCreateTestSuite) TestTransactionFindByUserIDFailureWithNotFoundCompany() {
	var (
		payload = models.TransactionCreatePayload{
			Company: "5f58f899b3d106cbfcecd333",
			Branch:  suite.data.Branch,
			User:    suite.data.User,
			Amount:  suite.data.Amount,
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *TransactionCreateTestSuite) TestTransactionFindByUserIDFailureWithNotFoundBranch() {
	var (
		payload = models.TransactionCreatePayload{
			Company: suite.data.Company,
			Branch:  "5f58f899b3d106cbfcecd444",
			User:    suite.data.User,
			Amount:  suite.data.Amount,
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *TransactionCreateTestSuite) TestTransactionFindByUserIDFailureWithNotFoundUser() {
	var (
		payload = models.TransactionCreatePayload{
			Company: suite.data.Company,
			Branch:  suite.data.Branch,
			User:    "5f58f899b3d106cbfcecd555",
			Amount:  suite.data.Amount,
		}
		response utils.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/transactions", utils.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionCreateTestSuite))
}

func setupDataTransaction() models.TransactionCreatePayload {
	payload := models.TransactionCreatePayload{
		Company: "5f58f899b3d106cbfcecd111",
		Branch:  "5f58f899b3d106cbfcecd112",
		User:    "5f58f899b3d106cbfcecd113",
		Amount:  50000,
	}
	return payload
}

func removeOldDataTransaction() {
	database.TransactionCol().DeleteMany(context.Background(), bson.M{})
}
