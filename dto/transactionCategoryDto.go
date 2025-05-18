package dto

type TransactionCategoryDto struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	UserID    string  `json:"user_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	CreatedBy string  `json:"created_by"`
	UpdatedBy *string `json:"updated_by"`
}

type CreateTransactionCategoryDto struct {
	Name   string `json:"name" binding:"required"`
}

type UpdateTransactionCategoryDto struct {
	Name  string `json:"name" binding:"required"`
}

