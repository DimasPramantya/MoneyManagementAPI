package domain

import (
	"time"

	"github.com/google/uuid"

	"github.com/dimas-pramantya/money-management/dto"
	. "github.com/dimas-pramantya/money-management/dto"
)

type TransactionCategory struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"modified_at"`
	UpdatedBy *string    `json:"updated_by" db:"modified_by"`
}

type TransactionCategoryRepository interface {
	FindByID(id int) (*TransactionCategory, error)
	FindByUserID(userID uuid.UUID) ([]TransactionCategory, error)
	Create(category *TransactionCategory) (*TransactionCategory, error)
	Update(category *TransactionCategory) (*TransactionCategory, error)
	Delete(id int) error
}

type TransactionCategoryUseCase interface {
	FindByID(id int) (*dto.TransactionCategoryDto, error)
	FindByUserID(userID uuid.UUID) ([]dto.TransactionCategoryDto, error)
	Create(req CreateTransactionCategoryDto, userID uuid.UUID) (*dto.TransactionCategoryDto, error)
	Update(req UpdateTransactionCategoryDto, id int, userID uuid.UUID) (*dto.TransactionCategoryDto, error)
	Delete(id int, userId uuid.UUID) error
}