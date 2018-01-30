package user

import (
	"fmt"
	"sync"

	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
	"golang.org/x/crypto/bcrypt"
)

type privateUserUsecase struct {
	userRepo    repository.UserRepository
	companyRepo repository.CompanyRepository
}

// GetAllUsers function to take care of get all users endpoints
func (u *privateUserUsecase) GetAllUsers() ([]*models.User, *models.HTTPError) {
	userList, err := u.userRepo.FetchAllUsers()
	if err != (*models.HTTPError)(nil) {
		return nil, err
	}

	wg := new(sync.WaitGroup)
	for i, usr := range userList {
		wg.Add(1)
		go func(i int, usr *models.User) {
			c, httpErr := u.companyRepo.GetCompanyByID(usr.CompanyID)
			if httpErr == (*models.HTTPError)(nil) {
				userList[i].Company = c
			}
			wg.Done()
		}(i, usr)
	}
	wg.Wait()

	return userList, nil
}

// CreateNewUser function to handle create new user business logics
func (u *privateUserUsecase) CreateNewUser(usr *models.User) *models.HTTPError {
	existing, httpErr := u.userRepo.GetUserByEmail(usr.Email)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	} else if existing != (*models.User)(nil) {
		return models.NewErrorUnprocessableEntity("Email Has Already Been Used For Another Account")
	}

	if len(usr.Password) < 10 {
		return models.NewErrorUnprocessableEntity("Password Needs to Be At Least 10 Letters")
	}

	existingComp, httpErr := u.companyRepo.GetCompanyByID(usr.CompanyID)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	} else if existingComp == (*models.Company)(nil) {
		return models.NewErrorUnprocessableEntity(fmt.Sprintf("Company with ID %d Doesnt Exist", usr.CompanyID))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 14)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	usr.Password = string(hash)

	usr, httpErr = u.userRepo.InsertNewUser(usr)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	}

	token, httpErr := u.userRepo.GenerateJWT(usr)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	}

	body := fmt.Sprintf("Your token is: %s", token)
	return u.userRepo.SendEmailToUser(usr, "First-time Login And Change Your Do-Do App Password", body)
}

// DeleteAnAccount delete an existing account
func (u *privateUserUsecase) DeleteAnAccount(id int64) *models.HTTPError {
	return u.userRepo.DeleteUser(id)
}

// Login login existing user
func (u *privateUserUsecase) Login(email, password string) (string, *models.HTTPError) {
	usr, httpErr := u.userRepo.GetUserByEmail(email)
	if httpErr != (*models.HTTPError)(nil) {
		return "", httpErr
	} else if usr == (*models.User)(nil) {
		return "", models.NewErrorUnauthorized("Invalid Email")
	}

	bcrErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if bcrErr != nil {
		return "", models.NewErrorUnauthorized("Invalid Password")
	}

	token, err := u.userRepo.GenerateJWT(usr)
	if err != (*models.HTTPError)(nil) {
		return "", models.NewErrorInternalServer(err.Error())
	}

	return token, nil
}

// ForgotPassword Trigger email with passsword reset link
func (u *privateUserUsecase) ForgotPassword(email string) *models.HTTPError {
	usr, httpErr := u.userRepo.GetUserByEmail(email)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	} else if usr == (*models.User)(nil) {
		return models.NewErrorNotFound("Email Is Not Registered")
	}

	token, httpErr := u.userRepo.GenerateJWT(usr)
	if httpErr != nil {
		return httpErr
	}

	emailBody := fmt.Sprintf("Your token is: %s", token)

	return u.userRepo.SendEmailToUser(usr, "Do-Do Password Reset Link", emailBody)
}

// ResetPassword reset password business logics
func (u *privateUserUsecase) ResetPassword(email, newPassword string) *models.HTTPError {
	usr, httpErr := u.userRepo.GetUserByEmail(email)
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	} else if usr == (*models.User)(nil) {
		return models.NewErrorNotFound("User Doesnt Exist")
	}

	if len(newPassword) < 10 {
		return models.NewErrorUnprocessableEntity("Password Needs to Be At Least 10 Letters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	usr.Password = string(hash)
	return u.userRepo.UpdateUser(usr)
}

// NewUserUsecase generate new user usecase
func NewUserUsecase(userRepo repository.UserRepository, companyRepo repository.CompanyRepository) usecase.UserUsecase {
	return &privateUserUsecase{
		userRepo,
		companyRepo,
	}
}
