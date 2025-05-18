package service

import (
	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/helper"
)

type TransactionSubCategoryService struct {
	transactionSubCategoryRepo domain.TransactionSubCategoryRepository
}

func (t *TransactionSubCategoryService) Create(req dto.CreateTransactionSubCategoryDto, userID string) (*dto.TransactionSubCategoryDto, error) {
	subCategory := &domain.TransactionSubCategory{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		CreatedBy:  userID,
	}

	createdSubCategory, err := t.transactionSubCategoryRepo.Create(subCategory)
	if err != nil {
		return nil, domain.InternalServerError("Failed to create transaction sub-category", err)
	}

	return mapTransactionSubCategoryToDto(createdSubCategory), nil
}

func (t *TransactionSubCategoryService) Delete(id int) error {
	subCategory, err := t.transactionSubCategoryRepo.FindByID(id)
	if err != nil {
		return domain.InternalServerError("Failed to find transaction sub-category", err)
	}
	if subCategory == nil {
		return domain.NotFoundError("Transaction sub-category not found", nil)
	}
	err = t.transactionSubCategoryRepo.Delete(id)
	if err != nil {
		return domain.InternalServerError("Failed to delete transaction sub-category", err)
	}
	return nil
}

func (t *TransactionSubCategoryService) FindAll() ([]dto.TransactionSubCategoryDto, error) {
	subCategories, err := t.transactionSubCategoryRepo.FindAll()
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction sub-categories", err)
	}

	var subCategoryDtos []dto.TransactionSubCategoryDto
	for _, subCategory := range subCategories {
		subCategoryDtos = append(subCategoryDtos, *mapTransactionSubCategoryToDto(&subCategory))
	}
	return subCategoryDtos, nil
}

func (t *TransactionSubCategoryService) FindByCategoryID(categoryID int) ([]dto.TransactionSubCategoryDto, error) {
	subCategories, err := t.transactionSubCategoryRepo.FindByCategoryID(categoryID)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction sub-categories by category ID", err)
	}

	var subCategoryDtos []dto.TransactionSubCategoryDto
	for _, subCategory := range subCategories {
		subCategoryDtos = append(subCategoryDtos, *mapTransactionSubCategoryToDto(&subCategory))
	}
	return subCategoryDtos, nil
}

func (t *TransactionSubCategoryService) FindByID(id int) (*dto.TransactionSubCategoryDto, error) {
	subCategory, err := t.transactionSubCategoryRepo.FindByID(id)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction sub-category", err)
	}
	if subCategory == nil {
		return nil, domain.NotFoundError("Transaction sub-category not found", nil)
	}
	return mapTransactionSubCategoryToDto(subCategory), nil
}

func (t *TransactionSubCategoryService) Update(req dto.UpdateTransactionSubCategoryDto, id int, userID string) (*dto.TransactionSubCategoryDto, error) {
	subCategory, err := t.transactionSubCategoryRepo.FindByID(id)
	if err != nil {
		return nil, domain.InternalServerError("Failed to find transaction sub-category", err)
	}
	if subCategory == nil {
		return nil, domain.NotFoundError("Transaction sub-category not found", nil)
	}

	subCategory.Name = req.Name
	subCategory.CategoryID = req.CategoryID
	subCategory.UpdatedBy = &userID

	updatedSubCategory, err := t.transactionSubCategoryRepo.Update(subCategory)
	if err != nil {
		return nil, domain.InternalServerError("Failed to update transaction sub-category", err)
	}

	return mapTransactionSubCategoryToDto(updatedSubCategory), nil
}

func NewTransactionSubCategoryService(transactionSubCategoryRepo domain.TransactionSubCategoryRepository) domain.TransactionSubCategoryUseCase {
	return &TransactionSubCategoryService{
		transactionSubCategoryRepo: transactionSubCategoryRepo,
	}
}

func mapTransactionSubCategoryToDto(subCategory *domain.TransactionSubCategory) *dto.TransactionSubCategoryDto {
	return &dto.TransactionSubCategoryDto{
		ID:         subCategory.ID,
		Name:       subCategory.Name,
		CategoryID: subCategory.CategoryID,
		CreatedBy:  subCategory.CreatedBy,
		UpdatedBy:  subCategory.UpdatedBy,
		CreatedAt:  *helper.TimeToString(&subCategory.CreatedAt),
		UpdatedAt:  helper.TimeToString(subCategory.UpdatedAt),
	}
}