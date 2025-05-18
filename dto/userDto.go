package dto

type RegisterDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`	
	Email   string `json:"email" binding:"required,email"`
}

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResUserDto struct {
	ID       string    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Balance  int64 `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy *string `json:"updated_by"`
}

type ResLoginDto struct {
	Token string `json:"token"`
	UserId string    `json:"user_id"`
}

type ReqUpdateUserDto struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type ReqUpdateUserPasswordDto struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ReqUpdateUserBalanceDto struct {
	Balance int64 `json:"balance" binding:"required"`
}