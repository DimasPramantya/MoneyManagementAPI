package domain

import (
	"time"

	"github.com/dimas-pramantya/money-management/dto"
)

type TransactionSubCategory struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	CategoryID int     `json:"category_id" db:"category_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy *string `json:"updated_by"`
}

type TransactionSubCategoryRepository interface {
	FindByID(id int) (*TransactionSubCategory, error)
	FindAll() ([]TransactionSubCategory, error)
	FindByCategoryID(categoryID int) ([]TransactionSubCategory, error)
	Create(subCategory *TransactionSubCategory) (*TransactionSubCategory, error)
	Update(subCategory *TransactionSubCategory) (*TransactionSubCategory, error)
	Delete(id int) error
}

type TransactionSubCategoryUseCase interface {
	FindByID(id int) (*dto.TransactionSubCategoryDto, error)
	FindAll() ([]dto.TransactionSubCategoryDto, error)
	FindByCategoryID(categoryID int) ([]dto.TransactionSubCategoryDto, error)
	Create(req dto.CreateTransactionSubCategoryDto, userID string) (*dto.TransactionSubCategoryDto, error)
	Update(req dto.UpdateTransactionSubCategoryDto, id int, userID string) (*dto.TransactionSubCategoryDto, error)
	Delete(id int) error
}