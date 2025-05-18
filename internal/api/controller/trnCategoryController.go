package controller

import (
	"strconv"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionCategoryController struct {
	TrnCategoryUC domain.TransactionCategoryUseCase
	validator *validation.Validator
}

func NewTransactionCategoryController(trnCategoryUC domain.TransactionCategoryUseCase, validator *validation.Validator) *TransactionCategoryController {
	return &TransactionCategoryController{
		TrnCategoryUC: trnCategoryUC,
		validator: validator,
	}
}

// CreateTransactionCategory godoc
// @Summary     Create Transaction Category
// @Description Create a new transaction category
// @Tags        category
// @Param       request body dto.CreateTransactionCategoryDto true "Create Transaction Category Payload"
// @Produce     json
// @Success     201 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories [POST]
func (uc *TransactionCategoryController) CreateTransactionCategory(ctx *gin.Context) {
	var req dto.CreateTransactionCategoryDto
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}
	userIDStr := ctx.MustGet("user_id").(string)
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	category, err := uc.TrnCategoryUC.Create(req, userUUID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, dto.BaseResponse{
		Message: "Transaction category created successfully",
		Data:    category,
		Code:    201,
	})
}

// UpdateTransactionCategory godoc
// @Summary     Update Transaction Category
// @Description Update an existing transaction category
// @Tags        category
// @Param       id   path int true "Transaction Category ID"
// @Param       request body dto.UpdateTransactionCategoryDto true "Update Transaction Category Payload"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/{id} [PUT]
func (uc *TransactionCategoryController) UpdateTransactionCategory(ctx *gin.Context) {
	var req dto.UpdateTransactionCategoryDto
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(domain.BadRequestError("Invalid ID", err))
		return
	}
	userIDStr := ctx.MustGet("user_id").(string)
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	category, err := uc.TrnCategoryUC.Update(req, idInt, userUUID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction category updated successfully",
		Data:    category,
		Code:    200,
	})
}

// DeleteTransactionCategory godoc
// @Summary     Delete Transaction Category
// @Description Delete a transaction category
// @Tags        category
// @Param       id path int true "Transaction Category ID"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/{id} [DELETE]
func (uc *TransactionCategoryController) DeleteTransactionCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(domain.BadRequestError("Invalid ID", err))
		return
	}
	userIDStr := ctx.MustGet("user_id").(string)
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = uc.TrnCategoryUC.Delete(idInt, userUUID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction category deleted successfully",
		Data:    nil,
		Code:    200,
	})
}

// GetTransactionCategoryByID godoc
// @Summary     Get Transaction Category by ID
// @Description Get a transaction category by ID
// @Tags        category
// @Param       id path int true "Transaction Category ID"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/{id} [GET]
func (uc *TransactionCategoryController) GetTransactionCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(domain.BadRequestError("Invalid ID", err))
		return
	}

	category, err := uc.TrnCategoryUC.FindByID(idInt)
	if err != nil {
		ctx.Error(err)
		return
	}
	if category == nil {
		ctx.JSON(404, dto.BaseResponse{
			Message: "Transaction category not found",
			Data:    nil,
			Code:    404,
		})
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction category retrieved successfully",
		Data:    category,
		Code:    200,
	})
}

// GetTransactionCategoriesByUserID godoc
// @Summary     Get Transaction Categories by User ID
// @Description Get all transaction categories for a user
// @Tags        category
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories [GET]
func (uc *TransactionCategoryController) GetTransactionCategoriesByUserID(ctx *gin.Context) {
	userIDStr := ctx.MustGet("user_id").(string)
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	categories, err := uc.TrnCategoryUC.FindByUserID(userUUID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction categories retrieved successfully",
		Data:    categories,
		Code:    200,
	})
}