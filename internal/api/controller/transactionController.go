package controller

import (

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionController struct {
	TransactionUseCase domain.TransactionUseCase
	validator          *validation.Validator
}

func NewTransactionController(transactionUseCase domain.TransactionUseCase, validator *validation.Validator) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
		validator:          validator,
	}
}

// CreateTransaction godoc
// @Summary     Create Transaction
// @Description Create a new transaction
// @Tags        transaction
// @Param       request body dto.CreateTransactionDto true "Create Transaction Payload"
// @Produce     json
// @Success     201 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transactions [POST]
func (uc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req dto.CreateTransactionDto
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

	transaction, err := uc.TransactionUseCase.Create(req, userUUID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, dto.BaseResponse{
		Message: "Transaction created successfully",
		Data:    transaction,
		Code:    201,
	})
}

// GetTransactionPaginated godoc
// @Summary     Get Transaction Paginated
// @Description Get transactions with pagination
// @Tags        transaction
// @Param       page query int false "Page number"
// @Param       limit query int false "Number of items per page"
// @Param       category_id query int false "Category ID"
// @Param       user_id query string false "User ID"
// @Param       sub_category_id query int false "Sub Category ID"
// @Param       start_date query string false "Start date (YYYY-MM-DD)"
// @Param       end_date query string false "End date (YYYY-MM-DD)"
// @Param       transaction_type query string false "Transaction type (income/expense)"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} domain.CustomError
// @Security    BearerAuth
// @Router      /transactions [GET]
func (uc *TransactionController) GetTransactionPaginated(ctx *gin.Context) {
	var req = dto.GetTransactionParams{
		CategoryID:   nil,
		UserId:	 nil,
		SubCategoryID: nil,
		StartDate:  nil,
		EndDate:    nil,
		TransactionType: nil,
		Limit:	1,
		Page: 	10,
	}

	userIDStr := ctx.MustGet("user_id").(string)

	req.UserId = &userIDStr

	transactions, err := uc.TransactionUseCase.FindByFilter(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dto.BaseResponse{
		Message: "Transactions retrieved successfully",
		Data:    transactions,
		Code:    200,
	})
}