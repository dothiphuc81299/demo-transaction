package apptest

import (
	"demo-transaction/modules/redis"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demo-transaction/config"
	"demo-transaction/modules/database"
	"demo-transaction/modules/zookeeper"
	"demo-transaction/routes"
	"demo-transaction/util"
)

func InitServer() *echo.Echo {
	config.InitENV()
	zookeeper.Connect()
	database.Connect()
	util.HelperConnect()
	redis.Connect()

	// New echo
	e := echo.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentLength, echo.HeaderContentType, echo.HeaderAuthorization},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		MaxAge:           600,
		AllowCredentials: false,
	}))

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	// Route
	routes.Boostrap(e)

	return e
}
