package controller

import (
	"fmt"
	"strconv"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionSubCategoryController struct {
	trnSubCategoryUseCase domain.TransactionSubCategoryUseCase
	trnCategoryUseCase   domain.TransactionCategoryUseCase
	validator *validation.Validator
}

func NewTransactionSubCategoryController(
	trnSubCategoryUseCase domain.TransactionSubCategoryUseCase, 
	trnCategoryUseCase domain.TransactionCategoryUseCase, 
	validator *validation.Validator,
) *TransactionSubCategoryController {
	return &TransactionSubCategoryController{
		trnSubCategoryUseCase: trnSubCategoryUseCase,
		trnCategoryUseCase:   trnCategoryUseCase,
		validator: validator,
	}
}

// CreateTransactionSubCategory godoc
// @Summary     Create Transaction SubCategory
// @Description Create a new transaction subcategory
// @Tags        category
// @Param       request body dto.CreateTransactionSubCategoryDto true "Create Transaction SubCategory Payload"
// @Produce     json
// @Success     201 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/sub-categories [POST]
func (uc *TransactionSubCategoryController) CreateTransactionSubCategory(ctx *gin.Context) {
	var req dto.CreateTransactionSubCategoryDto
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

	_, err = uc.trnCategoryUseCase.FindByID(req.CategoryID)
	if err != nil {
		ctx.Error(err)
		return
	}

	subCategory, err := uc.trnSubCategoryUseCase.Create(req, userUUID.String())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, dto.BaseResponse{
		Message: "Transaction subcategory created successfully",
		Data:    subCategory,
		Code:    201,
	})
}

// UpdateTransactionSubCategory godoc
// @Summary     Update Transaction SubCategory
// @Description Update an existing transaction subcategory
// @Tags        category
// @Param       id   path int true "Transaction SubCategory ID"
// @Param       request body dto.UpdateTransactionSubCategoryDto true "Update Transaction SubCategory Payload"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/sub-categories/{id} [PUT]
func (uc *TransactionSubCategoryController) UpdateTransactionSubCategory(ctx *gin.Context) {
	var req dto.UpdateTransactionSubCategoryDto
	err := uc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(domain.BadRequestError(fmt.Sprintf("Invalid transaction subcategory ID: %s", id), err))
		return
	}

	userIDStr := ctx.MustGet("user_id").(string)
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	subCategory, err := uc.trnSubCategoryUseCase.Update(req, idInt, userUUID.String())
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction subcategory updated successfully",
		Data:    subCategory,
		Code:    200,
	})
}

// DeleteTransactionSubCategory godoc
// @Summary     Delete Transaction SubCategory
// @Description Delete a transaction subcategory
// @Tags        category
// @Param       id path int true "Transaction SubCategory ID"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/sub-categories/{id} [DELETE]
func (uc *TransactionSubCategoryController) DeleteTransactionSubCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(domain.BadRequestError(fmt.Sprintf("Invalid transaction subcategory ID: %s", id), err))
		return
	}

	err = uc.trnSubCategoryUseCase.Delete(idInt)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction subcategory deleted successfully",
		Data:    nil,
		Code:    200,
	})
}

// FindAllTransactionSubCategories godoc
// @Summary     Find All Transaction SubCategories
// @Description Get all transaction subcategories
// @Tags        category
// @Produce     json
// @Param       categoryId query string false "Category ID (optional)"
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/sub-categories [GET]
func (uc *TransactionSubCategoryController) FindAllTransactionSubCategories(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")

	var (
		subCategories interface{}
		err           error
	)

	if categoryId != "" {
		categoryIdInt, err := strconv.Atoi(categoryId)
		if err != nil {
			ctx.Error(domain.BadRequestError(fmt.Sprintf("Invalid category ID: %s", categoryId), err))
			return
		}
		subCategories, err = uc.trnSubCategoryUseCase.FindByCategoryID(categoryIdInt)
	} else {
		subCategories, err = uc.trnSubCategoryUseCase.FindAll()
	}

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction subcategories retrieved successfully",
		Data:    subCategories,
		Code:    200,
	})
}

// FindTransactionSubCategoryByID godoc
// @Summary     Find Transaction SubCategory by ID
// @Description Get a transaction subcategory by ID
// @Tags        category
// @Param       id path int true "Transaction SubCategory ID"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transaction-categories/sub-categories/{id} [GET]
func (uc *TransactionSubCategoryController) FindTransactionSubCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(domain.BadRequestError(fmt.Sprintf("Invalid transaction subcategory ID: %s", id), err))
		return
	}

	subCategory, err := uc.trnSubCategoryUseCase.FindByID(idInt)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transaction subcategory retrieved successfully",
		Data:    subCategory,
		Code:    200,
	})
}