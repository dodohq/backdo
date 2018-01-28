package admin

import (
	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
	"golang.org/x/crypto/bcrypt"
)

type privateAdminUsecase struct {
	adminRepo repository.AdminRepository
}

// Login function to handle login of admin
// JWT on success
func (u *privateAdminUsecase) Login(username, password string) (string, *models.HTTPError) {
	admin, err := u.adminRepo.GetByEmail(username)
	if err != (*models.HTTPError)(nil) {
		return "", err
	}

	// check password
	bcrErr := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if bcrErr != nil {
		return "", models.NewErrorUnauthorized("Invalid Password")
	}

	token, err := u.adminRepo.GenerateJWT(admin)
	if err != (*models.HTTPError)(nil) {
		return "", models.NewErrorInternalServer(err.Error())
	}

	return token, nil
}

// NewAdminUsecase generate new admin usecase
func NewAdminUsecase(adminRepo repository.AdminRepository) usecase.AdminUsecase {
	return &privateAdminUsecase{
		adminRepo,
	}
}
