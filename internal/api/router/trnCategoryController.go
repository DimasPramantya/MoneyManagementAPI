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
	transactionSubCategoryRepo := pgrepository.NewTransactionSubCategoryPgRepository(db)

	// Usecases
	transactionCategoryUC := service.NewTransactionCategoryService(transactionCategoryRepo)
	transactionSubCategoryUC := service.NewTransactionSubCategoryService(transactionSubCategoryRepo)

	// Controllers
	transactionCategoryCtrl := controller.NewTransactionCategoryController(transactionCategoryUC, validator)
	transactionSubCategoryCtrl := controller.NewTransactionSubCategoryController(transactionSubCategoryUC, transactionCategoryUC, validator)

	// Routes
	rg.POST("", middleware.JwtMiddleware(), transactionCategoryCtrl.CreateTransactionCategory)
	rg.GET("", middleware.JwtMiddleware(), transactionCategoryCtrl.GetTransactionCategoriesByUserID)
	rg.PUT("/:id", middleware.JwtMiddleware(), transactionCategoryCtrl.UpdateTransactionCategory)

	rg.POST("/sub-categories", middleware.JwtMiddleware(), transactionSubCategoryCtrl.CreateTransactionSubCategory)
	rg.GET("/sub-categories", middleware.JwtMiddleware(), transactionSubCategoryCtrl.FindAllTransactionSubCategories)
	rg.GET("/sub-categories/:id", middleware.JwtMiddleware(), transactionSubCategoryCtrl.FindTransactionSubCategoryByID)
	rg.DELETE("/sub-categories/:id", middleware.JwtMiddleware(), transactionSubCategoryCtrl.DeleteTransactionSubCategory)
	rg.PUT("/sub-categories/:id", middleware.JwtMiddleware(), transactionSubCategoryCtrl.UpdateTransactionSubCategory)

	rg.GET("/:id", middleware.JwtMiddleware(), transactionCategoryCtrl.GetTransactionCategoryByID)
	rg.DELETE("/:id", middleware.JwtMiddleware(), transactionCategoryCtrl.DeleteTransactionCategory)
}