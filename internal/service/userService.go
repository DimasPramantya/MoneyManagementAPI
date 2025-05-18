package service

import (
	"fmt"

	"github.com/dimas-pramantya/money-management/dto"
	"github.com/dimas-pramantya/money-management/internal/api/middleware"
	"github.com/dimas-pramantya/money-management/internal/domain"
	"github.com/dimas-pramantya/money-management/utils/helper"
)

type UserService struct {
	userRepo domain.UserRepository
}

func (u *UserService) UpdateBalance(id string, req dto.ReqUpdateUserBalanceDto) (*dto.ResUserDto, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, domain.InternalServerError(fmt.Sprintf("Failed to find user with id %d", id), err)
	}

	if user == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("User with id %d not found.", id), nil)
	}

	user.Balance = req.Balance

	updatedUser, err := u.userRepo.UpdateBalance(user)
	if err != nil {
		return nil, domain.InternalServerError("Failed to update user balance", err)
	}

	return mapUserToResUserDto(updatedUser), nil
}

func (u *UserService) FindById(id string) (*dto.ResUserDto, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, domain.InternalServerError(fmt.Sprintf("Failed to find user with id %d", id), err)
	}

	if user == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("User with id %d not found.", id), nil)
	}

	return mapUserToResUserDto(user), nil
}

func (u *UserService) FindByUsername(username string) (*dto.ResUserDto, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return nil, domain.InternalServerError(fmt.Sprintf("Failed to find user with username %s", username), err)
	}

	if user == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("User with username %s not found.", username), nil)
	}

	res := &dto.ResUserDto{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}

	return res, nil
}

func (u *UserService) Login(req dto.LoginDto) (dto.ResLoginDto, error) {
	user, err := u.userRepo.FindByUsernameOrEmail(req.Username)
	if err != nil {
		return dto.ResLoginDto{}, domain.InternalServerError(fmt.Sprintf("Failed to find user with username %s", req.Username), err)
	}

	if user == nil {
		fmt.Println("User not found")
		return dto.ResLoginDto{}, domain.UnauthorizedError("Wrong username or password", nil)
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		fmt.Println("Wrong password")
		return dto.ResLoginDto{}, domain.UnauthorizedError("Wrong username or password", nil)
	}

	token, err := middleware.GenerateJwtToken(user.Username, user.ID.String())
	if err != nil {
		return dto.ResLoginDto{}, domain.InternalServerError("Failed to generate token", err)
	}

	res := dto.ResLoginDto{
		Token:  token,
		UserId: user.ID.String(),
	}

	return res, nil
}

func (u *UserService) Register(req dto.RegisterDto) (*dto.ResUserDto, error) {
	user, err := u.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, domain.InternalServerError(fmt.Sprintf("Failed to find user with username %s", req.Username), err)
	}

	if user != nil {
		return nil, domain.BadRequestError("Username already registered", nil)
	}

	user, err = u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, domain.InternalServerError(fmt.Sprintf("Failed to find user with email %s", req.Email), err)
	}
	if user != nil {
		return nil, domain.BadRequestError("Email already registered", nil)
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, domain.InternalServerError("Failed to hash password", err)
	}

	newUser := &domain.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err = u.userRepo.Create(newUser)
	if err != nil {
		return nil, domain.InternalServerError("Failed to create user", err)
	}

	return mapUserToResUserDto(user), nil
}

func (u *UserService) Update(id string, req dto.ReqUpdateUserDto) (*dto.ResUserDto, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, domain.InternalServerError(fmt.Sprintf("Failed to find user with id %d", id), err)
	}

	if user == nil {
		return nil, domain.NotFoundError(fmt.Sprintf("User with id %d not found.", id), nil)
	}

	user.Username = req.Username
	user.Email = req.Email

	updatedUser, err := u.userRepo.Update(user)
	if err != nil {
		return nil, domain.InternalServerError("Failed to update user", err)
	}

	return mapUserToResUserDto(updatedUser), nil
}

func (u *UserService) UpdatePassword(id string, req dto.ReqUpdateUserPasswordDto) error {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return domain.InternalServerError(fmt.Sprintf("Failed to find user with id %d", id), err)
	}

	if user == nil {
		return domain.NotFoundError(fmt.Sprintf("User with id %d not found.", id), nil)
	}

	if !helper.CheckPasswordHash(req.OldPassword, user.Password) {
		return domain.UnauthorizedError("Wrong old password", nil)
	}

	hashedPassword, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return domain.InternalServerError("Failed to hash password", err)
	}

	user.Password = hashedPassword
	user.UpdatedBy = &user.Username

	err = u.userRepo.UpdatePassword(user)
	if err != nil {
		return domain.InternalServerError("Failed to update password", err)
	}

	return nil
}

func NewUserService(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserService{
		userRepo: userRepo,
	}
}

func mapUserToResUserDto(user *domain.User) *dto.ResUserDto {
	return &dto.ResUserDto{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Balance:  user.Balance,
		CreatedAt: *helper.TimeToString(&user.CreatedAt),
		UpdatedAt: helper.TimeToString(user.UpdatedAt),
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
	}
}
