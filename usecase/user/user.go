package user

import (
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
)

type privateUserUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase generate new user usecase
func NewUserUsecase(userRepo repository.UserRepository) usecase.UserUsecase {
	return &privateUserUsecase{
		userRepo,
	}
}
