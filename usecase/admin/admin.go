package admin

import (
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
)

type privateAdminUsecase struct {
	adminRepo repository.AdminRepository
}

// Login function to handle login of admin
// JWT on success
func (u *privateAdminUsecase) Login(username, password string) (string, *models.HTTPError) {
	return "", models.NewErrorInternalServer("Not Implemented")
}

// NewAdminUsecase generate new admin usecase
func NewAdminUsecase(adminRepo repository.AdminRepository) usecase.AdminUsecase {
	return &privateAdminUsecase{
		adminRepo,
	}
}
