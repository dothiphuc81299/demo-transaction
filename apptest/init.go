package apptest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demo-transaction/config"
	"demo-transaction/modules/database"
	"demo-transaction/modules/redis"
	"demo-transaction/modules/zookeeper"
	"demo-transaction/routes"
	"demo-transaction/utils"
)

func InitServer() *echo.Echo {
	config.InitENV()
	zookeeper.Connect()
	database.Connect()
	utils.HelperConnect()
	redis.Connect()

	// New echo
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	// Route
	routes.Boostrap(e)
	return e
}
