package pgrepository

import (
	"database/sql"
	"fmt"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/google/uuid"
)

type transactionRepo struct {
	db *sql.DB
}

func (t *transactionRepo) CountByFilter(params dto.GetTransactionParams) (int, error) {
	query := `SELECT COUNT(*) FROM transactions WHERE user_id = $1`
	args := []interface{}{params.UserId}
	argPos := 2

	if params.CategoryID != nil {
		query += fmt.Sprintf(" AND category_id = $%d", argPos)
		args = append(args, *params.CategoryID)
		argPos++
	}
	if params.SubCategoryID != nil {
		query += fmt.Sprintf(" AND sub_category_id = $%d", argPos)
		args = append(args, *params.SubCategoryID)
		argPos++
	}
	if params.TransactionType != nil {
		query += fmt.Sprintf(" AND transaction_type = $%d", argPos)
		args = append(args, *params.TransactionType)
		argPos++
	}
	if params.StartDate != nil {
		query += fmt.Sprintf(" AND transaction_date >= $%d", argPos)
		args = append(args, *params.StartDate)
		argPos++
	}
	if params.EndDate != nil {
		query += fmt.Sprintf(" AND transaction_date <= $%d", argPos)
		args = append(args, *params.EndDate)
	}

	var count int
	err := t.db.QueryRow(query, args...).Scan(&count)
	return count, err
}

// Create implements domain.TransactionRepository.
func (t *transactionRepo) Create(tx *sql.Tx, transaction *domain.Transaction) (*domain.Transaction, error) {
	err := tx.QueryRow(`
		INSERT INTO transactions (ammount, transaction_category_id, transaction_sub_category_id, transaction_date, transaction_type, notes, user_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, ammount, transaction_category_id, transaction_sub_category_id, transaction_date, transaction_type, notes, user_id,
		created_at, created_by, updated_at, updated_by
	`, transaction.Ammount, transaction.CategoryID, transaction.SubCategoryID,
		transaction.TransactionDate, transaction.TransactionType,
		transaction.Notes, transaction.UserID.String(), transaction.CreatedBy).Scan(
		&transaction.ID,
		&transaction.Ammount,
		&transaction.CategoryID,
		&transaction.SubCategoryID,
		&transaction.TransactionDate,
		&transaction.TransactionType,
		&transaction.Notes,
		&transaction.UserID,
		&transaction.CreatedAt,
		&transaction.CreatedBy,
		&transaction.UpdatedAt,
		&transaction.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Delete implements domain.TransactionRepository.
func (t *transactionRepo) Delete(tx *sql.Tx, id int, userID uuid.UUID) error {
	_, err := tx.Exec(`DELETE FROM transactions WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepo) FindByFilter(params dto.GetTransactionParams) ([]dto.TransactionDto, error) {
	query := `
	SELECT 
		t.id, t.ammount, t.transaction_category_id, t.transaction_sub_category_id, t.transaction_date, 
		t.transaction_type, t.notes, t.user_id, t.created_at, t.created_by, 
		t.updated_at, t.updated_by,
		tc.name AS category,
		tsc.name AS sub_category
	FROM transactions t
	INNER JOIN transaction_categories tc ON t.transaction_category_id = tc.id
	LEFT JOIN transaction_sub_categories tsc ON t.transaction_sub_category_id = tsc.id
	WHERE t.user_id = $1
	`

	args := []interface{}{params.UserId}
	argPos := 2

	if params.CategoryID != nil {
		query += fmt.Sprintf(" AND t.category_id = $%d", argPos)
		args = append(args, *params.CategoryID)
		argPos++
	}
	if params.SubCategoryID != nil {
		query += fmt.Sprintf(" AND t.sub_category_id = $%d", argPos)
		args = append(args, *params.SubCategoryID)
		argPos++
	}
	if params.TransactionType != nil {
		query += fmt.Sprintf(" AND t.transaction_type = $%d", argPos)
		args = append(args, *params.TransactionType)
		argPos++
	}
	if params.StartDate != nil {
		query += fmt.Sprintf(" AND t.transaction_date >= $%d", argPos)
		args = append(args, *params.StartDate)
		argPos++
	}
	if params.EndDate != nil {
		query += fmt.Sprintf(" AND t.transaction_date <= $%d", argPos)
		args = append(args, *params.EndDate)
		argPos++
	}

	query += fmt.Sprintf(" ORDER BY t.transaction_date ASC LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, params.Limit, (params.Page-1)*params.Limit)

	rows, err := t.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []dto.TransactionDto
	for rows.Next() {
		var tx dto.TransactionDto
		if err := rows.Scan(
			&tx.ID, &tx.Ammount, &tx.CategoryID, &tx.SubCategoryID, &tx.TransactionDate,
			&tx.TransactionType, &tx.Notes, &tx.UserID, &tx.CreatedAt, &tx.CreatedBy,
			&tx.UpdatedAt, &tx.UpdatedBy, &tx.Category, &tx.SubCategory,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
}


func (t *transactionRepo) FindByID(id int) (*domain.Transaction, error) {
	row := t.db.QueryRow(`
		SELECT id, amount, category_id, sub_category_id, transaction_date, transaction_type, notes, user_id, created_at, created_by, updated_at, updated_by
		FROM transactions WHERE id = $1
	`, id)
	transaction := &domain.Transaction{}
	err := row.Scan(
		&transaction.ID,
		&transaction.Ammount,
		&transaction.CategoryID,
		&transaction.SubCategoryID,
		&transaction.TransactionDate,
		&transaction.TransactionType,
		&transaction.Notes,
		&transaction.UserID,
		&transaction.CreatedAt,
		&transaction.CreatedBy,
		&transaction.UpdatedAt,
		&transaction.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return transaction, nil
}

func (t *transactionRepo) Update(tx *sql.Tx, transaction *domain.Transaction) (*domain.Transaction, error) {
	err := tx.QueryRow(`
		UPDATE transactions
		SET ammount = $1, transaction_category_id = $2, transaction_sub_category_id = $3, transaction_date = $4, transaction_type = $5, notes = $6, updated_at = now(), updated_by = $7
		WHERE id = $8 RETURNING id, ammount, transaction_category_id, transaction_sub_category_id, transaction_date, transaction_type, notes, user_id, created_at, created_by,
		updated_at, updated_by
	`, transaction.Ammount, transaction.CategoryID, transaction.SubCategoryID,
		transaction.TransactionDate, transaction.TransactionType,
		transaction.Notes, transaction.UpdatedBy, transaction.ID).Scan(
		&transaction.ID,
		&transaction.Ammount,
		&transaction.CategoryID,
		&transaction.SubCategoryID,
		&transaction.TransactionDate,
		&transaction.TransactionType,
		&transaction.Notes,
		&transaction.UserID,
		&transaction.CreatedAt,
		&transaction.CreatedBy,
		&transaction.UpdatedAt,
		&transaction.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewTransactionRepo(db *sql.DB) domain.TransactionRepository {
	return &transactionRepo{db: db}
}
