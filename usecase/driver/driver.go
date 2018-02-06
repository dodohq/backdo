package driver

import (
	"fmt"

	"github.com/dodohq/backdo/models"
	"github.com/dodohq/backdo/repository"
	"github.com/dodohq/backdo/usecase"
)

type privateDriverUsecase struct {
	DriverRepo  repository.DriverRepository
	CompanyRepo repository.CompanyRepository
}

// NewDriverUsecase generate new driver usecase
func NewDriverUsecase(DriverRepo repository.DriverRepository, CompanyRepo repository.CompanyRepository) usecase.DriverUsecase {
	return &privateDriverUsecase{
		DriverRepo,
		CompanyRepo,
	}
}

// GetAllDriversOfCompany get all drivers under a company
func (u *privateDriverUsecase) GetAllDriversOfCompany(companyID int64) ([]*models.Driver, *models.HTTPError) {
	c, httpErr := u.CompanyRepo.GetCompanyByID(companyID)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	} else if c == (*models.Company)(nil) {
		return nil, models.NewErrorNotFound(fmt.Sprintf("Company with ID %d doesn't exist", companyID))
	}

	return u.DriverRepo.FetchDriversByCompany(companyID)
}

// GetDriverByID get driver by ID
func (u *privateDriverUsecase) GetDriverByID(id int64) (*models.Driver, *models.HTTPError) {
	d, httpErr := u.DriverRepo.FetchDriverByID(id)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	} else if d == (*models.Driver)(nil) {
		return nil, models.NewErrorNotFound(fmt.Sprintf("Driver with ID %d doesn't exist", id))
	}

	return u.DriverRepo.FetchDriverByID(id)
}

// GetDriverByPhoneNumber get driver profile with phone number
func (u *privateDriverUsecase) GetDriverByPhoneNumber(phoneNumber string) (*models.Driver, *models.HTTPError) {
	d, httpErr := u.DriverRepo.FetchDriverByPhoneNumber(phoneNumber)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	} else if d == (*models.Driver)(nil) {
		return nil, models.NewErrorNotFound(fmt.Sprintf("Driver with number %s doesn't exist", phoneNumber))
	}

	return d, nil
}

// CreateDriverProfile create new driver account
func (u *privateDriverUsecase) CreateDriverProfile(d *models.Driver) (*models.Driver, *models.HTTPError) {
	c, httpErr := u.CompanyRepo.GetCompanyByID(d.CompanyID)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	} else if c == (*models.Company)(nil) {
		return nil, models.NewErrorUnprocessableEntity(fmt.Sprintf("Company with ID %d doesn't exist", c.ID))
	}

	existing, httpErr := u.DriverRepo.FetchDriverByPhoneNumber(d.PhoneNumber)
	if httpErr != (*models.HTTPError)(nil) {
		return nil, httpErr
	} else if existing != (*models.Driver)(nil) {
		return nil, models.NewErrorUnprocessableEntity("Phone Number has been used for another account")
	}

	return u.DriverRepo.InsertNewDriver(d)
}

// RemoveDriverProfile remove driver profile
func (u *privateDriverUsecase) RemoveDriverProfile(id int64) *models.HTTPError {
	return u.DriverRepo.DeleteDriver(id)
}

// RequestForVerificationCode request for phone number verification code
func (u *privateDriverUsecase) RequestForVerificationCode(via string, countryCode int, phoneNumber string) *models.HTTPError {
	d, httpErr := u.DriverRepo.FetchDriverByPhoneNumber(fmt.Sprintf("+%d%s", countryCode, phoneNumber))
	if httpErr != (*models.HTTPError)(nil) {
		return httpErr
	} else if d == (*models.Driver)(nil) {
		return models.NewErrorNotFound(fmt.Sprintf("Driver with number +%d%s doesn't exist", countryCode, phoneNumber))
	}

	return u.DriverRepo.RequestVerificationCode(via, countryCode, phoneNumber)
}

// AuthenticateDriver authenticate driver, return JWT on success
func (u *privateDriverUsecase) AuthenticateDriver(countryCode int, phoneNumber, verificationCode string) (string, *models.HTTPError) {
	d, httpErr := u.DriverRepo.FetchDriverByPhoneNumber(fmt.Sprintf("+%d%s", countryCode, phoneNumber))
	if httpErr != (*models.HTTPError)(nil) {
		return "", httpErr
	} else if d == (*models.Driver)(nil) {
		return "", models.NewErrorNotFound(fmt.Sprintf("Driver with number +%d%s doesn't exist", countryCode, phoneNumber))
	}

	return u.DriverRepo.GenerateJWT(d)
}
