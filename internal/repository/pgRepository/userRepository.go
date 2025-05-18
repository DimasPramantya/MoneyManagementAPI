package pgrepository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/google/uuid"
)

type userPgRepository struct {
	db *sql.DB
}

func (u *userPgRepository) FindByUsernameOrEmail(username string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT id, username, email, password, created_at, created_by, updated_at, updated_by FROM users WHERE username = $1 OR email = $1`, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("Error scanning row:", err)
		return nil, err
	}
	return user, nil
}

func (u *userPgRepository) FindByEmail(email string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT id, username, email, password, created_at, created_by, updated_at, updated_by FROM users WHERE email = $1`, email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *userPgRepository) UpdatePassword(user *domain.User) error {
	_, err := u.db.Exec(`UPDATE users SET password = $1, updated_by = $2, updated_at = $3 WHERE id = $4`,
		user.Password, user.UpdatedBy, time.Now(), user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userPgRepository) Create(user *domain.User) (*domain.User, error) {
	err := u.db.QueryRow(`INSERT INTO users (id, username, password, email, created_by) VALUES ($1, $2, $3, $4, $5) Returning id, created_at, created_by`,
		uuid.New(), user.Username, user.Password, user.Email, "SYSTEM").Scan(&user.ID, &user.CreatedAt, &user.CreatedBy)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userPgRepository) FindById(id string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT id, username, email, password, created_at, created_by, updated_at, updated_by FROM users WHERE id = $1`, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *userPgRepository) FindByUsername(username string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT id, username, email, password, created_at, created_by, updated_at, updated_by FROM users WHERE username = $1`, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *userPgRepository) Update(user *domain.User) (*domain.User, error) {
	err := u.db.QueryRow(`UPDATE users SET username = $1, email = $2, updated_by = $3, updated_at = $4 WHERE id = $5 
		Returning updated_at, updated_by`,
		user.Username, user.Password, user.UpdatedBy, time.Now(), user.ID).Scan(&user.UpdatedAt, &user.UpdatedBy)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserPgRepository(db *sql.DB) domain.UserRepository {
	return &userPgRepository{db: db}
}
