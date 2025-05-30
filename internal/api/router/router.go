package router

import (
	"database/sql"

	"github.com/dimas-pramantya/money-management/utils/validation"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, db *sql.DB) {
	validator := validation.NewValidator()
	api := r.Group("/api")

	userRoute := api.Group("/users")
	InitUserRouter(userRoute, db, validator)

	transactionCategoryRoute := api.Group("/transaction-categories")
	InitCategoryRouter(transactionCategoryRoute, db, validator)

	transactionRoute := api.Group("/transactions")
	InitTransactionRouter(transactionRoute, db, validator)
}