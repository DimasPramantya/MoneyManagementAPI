package pgrepository

import (
	"database/sql"
	"time"

	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/google/uuid"
)

type transactionCategoryPgRepository struct {
	db *sql.DB
}

func (t *transactionCategoryPgRepository) Create(category *domain.TransactionCategory) (*domain.TransactionCategory, error) {
	err := t.db.QueryRow(`
		INSERT INTO transaction_categories (name, user_id, created_by)
		VALUES ($1, $2, $3) RETURNING id, name, user_id, created_at, created_by
	`, category.Name, category.UserID, category.CreatedBy).Scan(
		&category.ID, &category.Name, &category.UserID, &category.CreatedAt, &category.CreatedBy,
	)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (t *transactionCategoryPgRepository) Delete(id int) error {
	_, err := t.db.Exec(`DELETE FROM transaction_categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionCategoryPgRepository) FindByID(id int) (*domain.TransactionCategory, error) {
	row := t.db.QueryRow(`
		SELECT id, name, user_id, created_at, created_by, updated_at, updated_by
		FROM transaction_categories WHERE id = $1
	`, id)
	category := &domain.TransactionCategory{}
	err := row.Scan(
		&category.ID, &category.Name, &category.UserID, 
		&category.CreatedAt, &category.CreatedBy, &category.UpdatedAt, &category.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return category, nil
}

func (t *transactionCategoryPgRepository) FindByUserID(userID uuid.UUID) ([]domain.TransactionCategory, error) {
	rows, err := t.db.Query(`
		SELECT id, name, user_id, created_at, created_by, updated_at, updated_by
		FROM transaction_categories WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.TransactionCategory
	for rows.Next() {
		category := domain.TransactionCategory{}
		err := rows.Scan(
			&category.ID, &category.Name, &category.UserID,
			&category.CreatedAt, &category.CreatedBy, &category.UpdatedAt, &category.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (t *transactionCategoryPgRepository) Update(category *domain.TransactionCategory) (*domain.TransactionCategory, error) {
	err := t.db.QueryRow(`
		UPDATE transaction_categories
		SET name = $1, updated_by = $2, updated_at = $3
		WHERE id = $4 Returning type, name, updated_by, updated_at
	`,  category.Name, category.UpdatedBy, time.Now(), category.ID,
	).Scan(
		&category.Name, &category.UpdatedBy, &category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func NewTransactionCategoryPgRepository(db *sql.DB) domain.TransactionCategoryRepository {
	return &transactionCategoryPgRepository{db: db}
}
