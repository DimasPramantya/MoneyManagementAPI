package pgrepository

import (
	"database/sql"
	"time"

	"github.com/dimas-pramantya/money-management/internal/domain"
)

type transactionSubCategoryPgRepository struct {
	db *sql.DB
}

func (t *transactionSubCategoryPgRepository) Create(subCategory *domain.TransactionSubCategory) (*domain.TransactionSubCategory, error) {
	err := t.db.QueryRow(`
		INSERT INTO transaction_sub_categories (name, transaction_category_id, created_by)
		VALUES ($1, $2, $3) RETURNING id, name, transaction_category_id, created_at, created_by
	`, subCategory.Name, subCategory.CategoryID, subCategory.CreatedBy).Scan(
		&subCategory.ID, &subCategory.Name, &subCategory.CategoryID,
		&subCategory.CreatedAt, &subCategory.CreatedBy,
	)
	if err != nil {
		return nil, err
	}
	return subCategory, nil
}

func (t *transactionSubCategoryPgRepository) Delete(id int) error {
	_, err := t.db.Exec(`DELETE FROM transaction_sub_categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionSubCategoryPgRepository) FindAll() ([]domain.TransactionSubCategory, error) {
	rows, err := t.db.Query(`
		SELECT id, name, transaction_category_id, created_at, created_by, updated_at, updated_by
		FROM transaction_sub_categories
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCategories []domain.TransactionSubCategory
	for rows.Next() {
		var subCategory domain.TransactionSubCategory
		err := rows.Scan(
			&subCategory.ID, &subCategory.Name, &subCategory.CategoryID,
			&subCategory.CreatedAt, &subCategory.CreatedBy,
			&subCategory.UpdatedAt, &subCategory.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		subCategories = append(subCategories, subCategory)
	}
	return subCategories, nil
}

func (t *transactionSubCategoryPgRepository) FindByCategoryID(categoryID int) ([]domain.TransactionSubCategory, error) {
	rows, err := t.db.Query(`
		SELECT id, name, transaction_category_id, created_at, created_by, updated_at, updated_by
		FROM transaction_sub_categories WHERE transaction_category_id = $1
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCategories []domain.TransactionSubCategory
	for rows.Next() {
		var subCategory domain.TransactionSubCategory
		err := rows.Scan(
			&subCategory.ID, &subCategory.Name, &subCategory.CategoryID,
			&subCategory.CreatedAt, &subCategory.CreatedBy,
			&subCategory.UpdatedAt, &subCategory.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		subCategories = append(subCategories, subCategory)
	}
	return subCategories, nil
}

func (t *transactionSubCategoryPgRepository) FindByID(id int) (*domain.TransactionSubCategory, error) {
	row := t.db.QueryRow(`
		SELECT id, name, transaction_category_id, created_at, created_by, updated_at, updated_by
		FROM transaction_sub_categories WHERE id = $1
	`, id)
	subCategory := &domain.TransactionSubCategory{}
	err := row.Scan(
		&subCategory.ID, &subCategory.Name, &subCategory.CategoryID,
		&subCategory.CreatedAt, &subCategory.CreatedBy,
		&subCategory.UpdatedAt, &subCategory.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return subCategory, nil
}

func (t *transactionSubCategoryPgRepository) Update(subCategory *domain.TransactionSubCategory) (*domain.TransactionSubCategory, error) {
	err := t.db.QueryRow(`
		UPDATE transaction_sub_categories SET name = $1, transaction_category_id = $2, updated_by = $3, updated_at = $4
		WHERE id = $5 RETURNING id, name, transaction_category_id, created_at, created_by, updated_at, updated_by
	`, subCategory.Name, subCategory.CategoryID, subCategory.CreatedBy, time.Now(), subCategory.UpdatedBy,
	).Scan(
		&subCategory.ID, &subCategory.Name, &subCategory.CategoryID,
		&subCategory.CreatedAt, &subCategory.CreatedBy,
		&subCategory.UpdatedAt, &subCategory.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return subCategory, nil
}

func NewTransactionSubCategoryPgRepository(db *sql.DB) domain.TransactionSubCategoryRepository {
	return &transactionSubCategoryPgRepository{db: db}
}
