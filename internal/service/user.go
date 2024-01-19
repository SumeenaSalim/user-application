package service

import (
	"context"
	"user-app/internal/dto/output"
	"user-app/internal/repository"
	"user-app/internal/utils"
)

// UserServiceInterface defines the contract for the user service
type UserServiceInterface interface {
	CreateUser(ctx context.Context) (output.UserResponse, error)
	UpdateUser(ctx context.Context) (output.UserResponse, error)
}

// UserService is responsible for handling user-related business logic
type UserService struct {
	userServiceRepo repository.UserRepositoryInterface
}

// NewUserService creates a new instance of UserService
func NewUserService(userServiceRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userServiceRepo: userServiceRepo,
	}
}

// CreateUser creates a new user and returns the response
func (u *UserService) CreateUser(ctx context.Context) (output.UserResponse, error) {

	request := utils.GetCreateUserRequestFromCtx(ctx)

	response := output.UserResponse{}

	userData, err := u.userServiceRepo.CreateUser(ctx, request)
	if err != nil {
		return response, err
	}
	response.Id = (userData.ID).String()
	response.Email = userData.Email
	response.Name = userData.Name
	return response, nil
}

func (u *UserService) UpdateUser(ctx context.Context) (output.UserResponse, error) {
	request := utils.GetUpdateUserRequestFromCtx(ctx)
	userId := utils.GetUserIDFromCtx(ctx)

	response := output.UserResponse{}

	userData, err := u.userServiceRepo.UpdateUser(ctx, userId, request)
	if err != nil {
		return response, err
	}
	response.Id = (userData.ID).String()
	response.Email = userData.Email
	response.Name = userData.Name
	return response, nil
}