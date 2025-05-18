package service

import (
	"fmt"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/helper"
	"github.com/google/uuid"
)

type TransactionCategoryService struct {
	transactionCategoryRepo domain.TransactionCategoryRepository
}

func (t *TransactionCategoryService) Create(req dto.CreateTransactionCategoryDto, userID uuid.UUID) (*dto.TransactionCategoryDto, error) {
	category := &domain.TransactionCategory{
		Name:   req.Name,
		UserID: userID,
		CreatedBy: userID.String(),
	}

	createdCategory, err := t.transactionCategoryRepo.Create(category)
	if err != nil {
		return nil, domain.InternalServerError("Failed to create transaction category", err)
	}

	return mapTransactionCategoryToDto(createdCategory), nil
}

func (t *TransactionCategoryService) Delete(id int, userId uuid.UUID) error {
	category, err := t.transactionCategoryRepo.FindByID(id)
	if err != nil {
		return domain.InternalServerError("Failed to find transaction category", err)
	}
	if category == nil {
		return domain.NotFoundError(fmt.Sprintf("Transaction category with id %d not found", id), nil)
	}
	if category.UserID != userId {
		return domain.UnauthorizedError("Unauthorized", nil)
	}
	err = t.transactionCategoryRepo.Delete(id)
	if err != nil {
		return domain.InternalServerError("Failed to delete transaction category", err)
	}
	return nil
}

func (t *TransactionCategoryService) FindByID(id int) (*dto.TransactionCategoryDto, error) {
	category, err := t.transactionCategoryRepo.FindByID(id)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction category", err)
	}
	if category == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("Transaction category with id %d not found", id), nil)
	}
	return mapTransactionCategoryToDto(category), nil
}

func (t *TransactionCategoryService) FindByUserID(userID uuid.UUID) ([]dto.TransactionCategoryDto, error) {
	categories, err := t.transactionCategoryRepo.FindByUserID(userID)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction categories", err)
	}
	result := make([]dto.TransactionCategoryDto, len(categories))
	for i, category := range categories {
		result[i] = *mapTransactionCategoryToDto(&category)
	}
	return result, nil
}

func (t *TransactionCategoryService) Update(req dto.UpdateTransactionCategoryDto, id int, userID uuid.UUID) (*dto.TransactionCategoryDto, error) {
	category, err := t.transactionCategoryRepo.FindByID(id)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction category", err)
	}
	if category == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("Transaction category with id %d not found", id), nil)
	}

	category.Name = req.Name
	category.UserID = userID
	updatedBy := userID.String()
	category.UpdatedBy = &updatedBy

	updatedCategory, err := t.transactionCategoryRepo.Update(category)
	if err != nil {
		return nil, domain.InternalServerError("Failed to update transaction category", err)
	}

	return mapTransactionCategoryToDto(updatedCategory), nil
}

func NewTransactionCategoryService(transactionCategoryRepo domain.TransactionCategoryRepository) domain.TransactionCategoryUseCase {
	return &TransactionCategoryService{
		transactionCategoryRepo: transactionCategoryRepo,
	}
}

func mapTransactionCategoryToDto(category *domain.TransactionCategory) *dto.TransactionCategoryDto {
	return &dto.TransactionCategoryDto{
		ID:        category.ID,
		Name:      category.Name,
		UserID:    category.UserID.String(),
		CreatedAt: *helper.TimeToString(&category.CreatedAt),
		UpdatedAt: helper.TimeToString(category.UpdatedAt),
		CreatedBy: category.CreatedBy,
		UpdatedBy: category.UpdatedBy,
	}
}