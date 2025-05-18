package controller

import (
	"net/http"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/validation"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUC    domain.UserUsecase
	validator *validation.Validator
}

func NewUserController(userUC domain.UserUsecase, validator *validation.Validator) *UserController {
	return &UserController{
		UserUC:    userUC,
		validator: validator,
	}
}

// Register godoc
// @Summary     Register
// @Description Register new user
// @Tags        auth
// @Param       request body dto.RegisterDto true "Register Payload"
// @Produce     json
// @Success     201 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Router      /users/register [POST]
func (uc *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterDto
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	user, err := uc.UserUC.Register(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.BaseResponse{
		Message: "User registered successfully",
		Data:    user,
		Code:    http.StatusCreated,
	})
}

// Login godoc
// @Summary     Login
// @Description Login user
// @Tags        auth
// @Param       request body dto.LoginDto true "Login Payload"
// @Produce     json
// @Success     200 {object} dto.ResLoginDto
// @Failure     400 {object} domain.CustomError
// @Router      /users/login [POST]
func (uc *UserController) Login(ctx *gin.Context) {
	var req dto.LoginDto
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	user, err := uc.UserUC.Login(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.BaseResponse{
		Message: "User logged in successfully",
		Data:    user,
		Code:    http.StatusOK,
	})
}

// GetUserProfile godoc
// @Summary     Get User Profile By JWT
// @Description Get user profile by JWT
// @Tags        users
// @Produce     json
// @Success     200 {array} dto.ResUserDto
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /users/profile [get]
func (uc *UserController) GetUserProfile(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	username, _ := ctx.Get("username")
	if userID == nil || username == nil {
		err := domain.UnauthorizedError("Unauthorized", nil)
		ctx.Error(err)
		return
	}

	userIDStr, _ := userID.(string)
	user, err := uc.UserUC.FindById(userIDStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.BaseResponse{
		Message: "Get User Profile Success",
		Data:    user,
		Code:    http.StatusOK,
	})
}

// UpdateUserBalance godoc
// @Summary     Update User Balance
// @Description Update user balance
// @Tags        users
// @Param       request body dto.ReqUpdateUserBalanceDto true "Update User Balance Payload"
// @Produce     json
// @Success     200 {object} dto.ResUserDto
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /users/balance [patch]
func (uc *UserController) UpdateUserBalance(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	if userID == nil {
		err := domain.UnauthorizedError("Unauthorized", nil)
		ctx.Error(err)
		return
	}

	var req dto.ReqUpdateUserBalanceDto
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	userIDStr, _ := userID.(string)
	user, err := uc.UserUC.UpdateBalance(userIDStr, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.BaseResponse{
		Message: "User balance updated successfully",
		Data:    user,
		Code:    http.StatusOK,
	})
}