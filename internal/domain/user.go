package domain

import (
	"database/sql"
	"time"

	. "github.com/dimas-pramantya/money-management/dto"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string `json:"username"` 
	Password string `json:"password"`
	Email    string `json:"email"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy *string `json:"updated_by"`
}

type UserRepository interface {
	FindById(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
	UpdatePassword(user *User) (error)
	FindByUsernameOrEmail(username string) (*User, error)
	UpdateBalance(user *User) (*User, error)
	UpdateBalanceTx(tx *sql.Tx, user *User) (*User, error)
}

type UserUsecase interface {
	FindById(id string) (*ResUserDto, error)
	FindByUsername(username string) (*ResUserDto, error)
	Login(req LoginDto) (ResLoginDto, error)
	Register(req RegisterDto) (*ResUserDto, error)
	Update(id string, req ReqUpdateUserDto) (*ResUserDto, error)
	UpdatePassword(id string, req ReqUpdateUserPasswordDto) (error)
	UpdateBalance(id string, req ReqUpdateUserBalanceDto) (*ResUserDto, error)
}