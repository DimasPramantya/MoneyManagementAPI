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

func InitTransactionRouter(rg *gin.RouterGroup, db *sql.DB, validator *validation.Validator) {
	// Repositories
	transactionCategoryRepo := pgrepository.NewTransactionCategoryPgRepository(db)
	transactionSubCategoryRepo := pgrepository.NewTransactionSubCategoryPgRepository(db)
	transactionRepo := pgrepository.NewTransactionRepo(db)
	userRepo := pgrepository.NewUserPgRepository(db)

	// Usecases
	transactionUseCase := service.NewTransactionService(transactionRepo, transactionCategoryRepo, transactionSubCategoryRepo, userRepo, db)

	// Controllers
	transactionController := controller.NewTransactionController(transactionUseCase, validator)

	// Routes
	rg.POST("", middleware.JwtMiddleware(), transactionController.CreateTransaction)
	rg.GET("", middleware.JwtMiddleware(), transactionController.GetTransactionPaginated)
}