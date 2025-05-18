package dto

type TransactionSubCategoryDto struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"category_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
	CreatedBy  string  `json:"created_by"`
	UpdatedBy  *string `json:"updated_by"`
}

type CreateTransactionSubCategoryDto struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type UpdateTransactionSubCategoryDto struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}
