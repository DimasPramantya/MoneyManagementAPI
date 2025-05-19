package service

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/helper"
	"github.com/google/uuid"
)

type TransactionService struct {
	transactionRepo    domain.TransactionRepository
	trnCategoryRepo    domain.TransactionCategoryRepository
	trnSubCategoryRepo domain.TransactionSubCategoryRepository
	userRepo           domain.UserRepository
	db 				   *sql.DB
}

func (t *TransactionService) Create(req dto.CreateTransactionDto, userID uuid.UUID) (*dto.TransactionDto, error) {
	category, err := t.trnCategoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction category", err)
	}
	if category == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("Category with id %d not found", req.CategoryID), nil)
	}

	subCategory := &domain.TransactionSubCategory{}
	if req.SubCategoryID != nil {
		subCategory, err = t.trnSubCategoryRepo.FindByID(*req.SubCategoryID)
		if err != nil {
			return nil, domain.InternalServerError("Failed to find transaction sub-category", err)
		}
		if subCategory == nil {
			return nil, domain.NotFoundError(fmt.Sprintf("Sub-category with id %d not found", *req.SubCategoryID), nil)
		}
	}

	tx, err := t.db.Begin()

	user, err := t.userRepo.FindById(userID.String())
	if err != nil {
		return nil, domain.InternalServerError("Failed to find user", err)
	}
	if user == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("User with id %s not found", userID), nil)
	}
	ammount := req.Amount
	if req.TransactionType == "income" {
		ammount = req.Amount
	}else if req.TransactionType == "expense" {
		ammount = -req.Amount
	}else{
		return nil, domain.BadRequestError("Invalid transaction type", nil)
	}
	user.Balance += ammount
	t.userRepo.UpdateBalanceTx(tx, user)

	transactionDate, err := helper.StringToDate(req.TransactionDate)
	if err != nil {
		return nil, domain.BadRequestError("Invalid transaction date format", err)
	}
	transaction := &domain.Transaction{
		Ammount:          req.Amount,
		CategoryID:      req.CategoryID,
		SubCategoryID:   req.SubCategoryID,
		TransactionDate: *transactionDate,
		TransactionType: req.TransactionType,
		Notes:            req.Note,
		UserID:          userID,
	}

	createdTransaction, err := t.transactionRepo.Create(tx, transaction)
	if err != nil {
		return nil, domain.InternalServerError("Failed to create transaction", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, domain.InternalServerError("Failed to commit transaction", err)
	}

	return mapTransactionToDto(createdTransaction, category.Name, &subCategory.Name), nil
}

// Delete implements domain.TransactionUseCase.
func (t *TransactionService) Delete(id int, userID uuid.UUID) error {
	panic("unimplemented")
}

func (t *TransactionService) FindByFilter(params dto.GetTransactionParams) (dto.PaginationResponse[dto.TransactionDto], error) {
	total, err := t.transactionRepo.CountByFilter(params)
	if err != nil {
		return dto.PaginationResponse[dto.TransactionDto]{}, domain.InternalServerError("Failed to count transactions", err)
	}

	records, err := t.transactionRepo.FindByFilter(params)
	if err != nil {
		return dto.PaginationResponse[dto.TransactionDto]{}, domain.InternalServerError("Failed to fetch transactions", err)
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.Limit)))
	var nextPage *int
	if params.Page < totalPages {
		next := params.Page + 1
		nextPage = &next
	}
	var prevPage *int
	if params.Page > 1 {
		prev := params.Page - 1
		prevPage = &prev
	}

	return dto.PaginationResponse[dto.TransactionDto]{
		TotalRecords:  total,
		TotalPages:    totalPages,
		CurrentPage:   params.Page,
		Limit:         params.Limit,
		Records:       records,
		NextPage:      nextPage,
		PreviousPage:  prevPage,
	}, nil
}

// FindByID implements domain.TransactionUseCase.
func (t *TransactionService) FindByID(id int) (*dto.TransactionDto, error) {
	panic("unimplemented")
}

// Update implements domain.TransactionUseCase.
func (t *TransactionService) Update(req dto.UpdateTransactionDto, id int, userID uuid.UUID) (*dto.TransactionDto, error) {
	panic("unimplemented")
}

func NewTransactionService(
	transactionRepo domain.TransactionRepository,
	trnCategoryRepo domain.TransactionCategoryRepository,
	trnSubCategoryRepo domain.TransactionSubCategoryRepository,
	userRepo domain.UserRepository,
	db *sql.DB,
) domain.TransactionUseCase {
	return &TransactionService{
		transactionRepo:    transactionRepo,
		trnCategoryRepo:    trnCategoryRepo,
		trnSubCategoryRepo: trnSubCategoryRepo,
		userRepo:           userRepo,
		db: 				db,
	}
}

func mapTransactionToDto(transaction *domain.Transaction, category string, subCategory *string) *dto.TransactionDto {
	return &dto.TransactionDto{
		ID:              transaction.ID,
		Ammount:          transaction.Ammount,
		CategoryID:      transaction.CategoryID,
		SubCategoryID:   transaction.SubCategoryID,
		TransactionDate: *helper.DateToString(&transaction.TransactionDate),
		TransactionType: transaction.TransactionType,
		Notes:            transaction.Notes,
		UserID:          transaction.UserID.String(),
		CreatedAt:       *helper.TimeToString(&transaction.CreatedAt),
		UpdatedAt:       helper.TimeToString(transaction.UpdatedAt),
		CreatedBy:       transaction.CreatedBy,
		UpdatedBy:       transaction.UpdatedBy,
	}
}
