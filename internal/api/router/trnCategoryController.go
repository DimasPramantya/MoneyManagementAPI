package router

import (
	"database/sql"

	"github.com/dimas-pramantya/money-management/internal/api/controller"
	"github.com/dimas-pramantya/money-management/internal/api/middleware"
	pgrepository "github.com/dimas-pramantya/money-management/internal/repository/pgRepository"
	"github.com/dimas-pramantya/money-management/internal/service"
	"github.com/dimas-pramantya/money-management/utils/validation"
	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(rg *gin.RouterGroup, db *sql.DB, validator *validation.Validator) {
	// Repositories
	transactionCategoryRepo := pgrepository.NewTransactionCategoryPgRepository(db)

	// Usecases
	transactionCategoryUC := service.NewTransactionCategoryService(transactionCategoryRepo)

	// Controllers
	transactionCategoryCtrl := controller.NewTransactionCategoryController(transactionCategoryUC, validator)

	// Routes
	rg.POST("", middleware.JwtMiddleware(), transactionCategoryCtrl.CreateTransactionCategory)
	rg.GET("", middleware.JwtMiddleware(), transactionCategoryCtrl.GetTransactionCategoriesByUserID)
	rg.GET("/:id", middleware.JwtMiddleware(), transactionCategoryCtrl.GetTransactionCategoryByID)
	rg.DELETE("/:id", middleware.JwtMiddleware(), transactionCategoryCtrl.DeleteTransactionCategory)
	rg.PUT("/:id", middleware.JwtMiddleware(), transactionCategoryCtrl.UpdateTransactionCategory)
}