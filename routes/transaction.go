package routes

import (
	"github.com/labstack/echo/v4"

	"demo-transaction/controllers"
	"demo-transaction/validations"
)

//Transaction func ...
func Transaction(e *echo.Echo) {
	routes := e.Group("/transactions")

	routes.POST("", controllers.TransactionCreate, validations.TransactionCreate, companyCheckExistedByID, branchCheckExistedByID, userCheckExistedByID)
}
