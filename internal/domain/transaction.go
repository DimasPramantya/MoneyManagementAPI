package domain

import (
	"database/sql"
	"time"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/google/uuid"
)

type Transaction struct {
	ID              int   `json:"id"`
	Ammount          int64 `json:"amount"`
	CategoryID      int   `json:"category_id"`
	SubCategoryID   *int  `json:"sub_category_id"`
	TransactionDate time.Time  `json:"transaction_date"`
	TransactionType string `json:"transaction_type"`
	Notes            *string `json:"note"`
	UserID         	uuid.UUID `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	CreatedBy       string `json:"created_by"`
	UpdatedBy       *string `json:"updated_by"`
}

type TransactionRepository interface {
	FindByID(id int) (*Transaction, error)
	FindByFilter(params dto.GetTransactionParams) ([]dto.TransactionDto, error)
	CountByFilter(params dto.GetTransactionParams) (int, error)
	Create(tx *sql.Tx, transaction *Transaction) (*Transaction, error)
	Update(tx *sql.Tx, transaction *Transaction) (*Transaction, error)
	Delete(tx *sql.Tx, id int, userID uuid.UUID) error
}

type TransactionUseCase interface {
	FindByID(id int) (*dto.TransactionDto, error)
	FindByFilter(params dto.GetTransactionParams) (dto.PaginationResponse[dto.TransactionDto], error)
	Create(req dto.CreateTransactionDto, userID uuid.UUID) (*dto.TransactionDto, error)
	Update(req dto.UpdateTransactionDto, id int, userID uuid.UUID) (*dto.TransactionDto, error)
	Delete(id int, userID uuid.UUID) error
}
