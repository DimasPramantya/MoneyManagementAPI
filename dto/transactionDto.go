package dto

type TransactionDto struct {
	ID         int     `json:"id"`
	Ammount     int64 `json:"amount"`
	CategoryID int     `json:"category_id"`
	Category   string  `json:"category"`
	SubCategoryID *int  `json:"sub_category_id"`
	SubCategory   *string `json:"sub_category"`
	TransactionDate string `json:"transaction_date"`
	TransactionType string `json:"transaction_type"`
	Notes        *string `json:"note"`
	UserID      string  `json:"user_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	CreatedBy   string  `json:"created_by"`
	UpdatedBy   *string `json:"updated_by"`
}

type GetTransactionParams struct {
	CategoryID    *int    `form:"category_id"`
	UserId 	 *string 	  `form:"user_id"`
	SubCategoryID *int    `form:"sub_category_id"`
	StartDate     *string `form:"start_date"`
	EndDate       *string `form:"end_date"`
	TransactionType *string `form:"transaction_type"`
	Limit         int    `form:"limit"`
	Page          int    `form:"page"`
}

type CreateTransactionDto struct {
	Amount          int64 `json:"amount" binding:"required"`
	CategoryID      int   `json:"category_id" binding:"required"`
	SubCategoryID   *int  `json:"sub_category_id"`
	TransactionDate string `json:"transaction_date" binding:"required"`
	TransactionType string `json:"transaction_type" binding:"required"`
	Note            *string `json:"note"`
}

type UpdateTransactionDto struct {
	Amount          int64 `json:"amount" binding:"required"`
	CategoryID      int   `json:"category_id" binding:"required"`
	SubCategoryID   *int  `json:"sub_category_id"`
	TransactionDate string `json:"transaction_date" binding:"required"`
	TransactionType string `json:"transaction_type" binding:"required"`
	Note            *string `json:"note"`
}
