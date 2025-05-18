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

func InitUserRouter(rg *gin.RouterGroup, db *sql.DB, validator *validation.Validator) {
	// Repositories
	userRepo := pgrepository.NewUserPgRepository(db)

	// Usecases
	userUC := service.NewUserService(userRepo)

	// Controllers
	userCtrl := controller.NewUserController(userUC, validator)

	// Routes
	rg.POST("/register", userCtrl.Register)
	rg.POST("/login", userCtrl.Login)
	rg.GET("/profile", middleware.JwtMiddleware(),  userCtrl.GetUserProfile)
	rg.PATCH("/balance", middleware.JwtMiddleware(), userCtrl.UpdateUserBalance)
}