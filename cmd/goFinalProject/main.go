// @title       Money Management API
// @version     1.0
// @description API for money management, to track expense and income \n\nTo authorize, click "Authorize" and enter your JWT token in this format:\n**Bearer &lt;your_token&gt;**
// @BasePath    /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"github.com/dimas-pramantya/money-management/internal/api/router"
	. "github.com/dimas-pramantya/money-management/internal/api/middleware"
	"github.com/dimas-pramantya/money-management/internal/configs"
	"github.com/dimas-pramantya/money-management/internal/database/connection"
	"github.com/dimas-pramantya/money-management/internal/database/migration"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	_ "github.com/dimas-pramantya/money-management/docs"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	migration.Initiator(connection.DBConnections)
	defer connection.DBConnections.Close()

	r := gin.Default()
	r.Use(GlobalExceptionHandler())

	router.Init(r, connection.DBConnections)

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}