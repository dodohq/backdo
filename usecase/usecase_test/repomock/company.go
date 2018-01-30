package repomock

import (
	"github.com/dodohq/backdo/models"
	"github.com/stretchr/testify/mock"
)

// CompanyRepoMock mock of company repo
type CompanyRepoMock struct {
	mock.Mock
}

// GetAllCompany mock of function
func (m *CompanyRepoMock) GetAllCompany() ([]*models.Company, *models.HTTPError) {
	args := m.Called()

	return args[0].([]*models.Company), args[1].(*models.HTTPError)
}

// GetCompanyByID mock of func
func (m *CompanyRepoMock) GetCompanyByID(id int64) (*models.Company, *models.HTTPError) {
	args := m.Called(id)

	return args[0].(*models.Company), args[1].(*models.HTTPError)
}

// InsertNewCompany mock of func
func (m *CompanyRepoMock) InsertNewCompany(c *models.Company) (*models.Company, *models.HTTPError) {
	args := m.Called(c)

	return args[0].(*models.Company), args[1].(*models.HTTPError)
}

// DeleteACompany mock of func
func (m *CompanyRepoMock) DeleteACompany(id int64) (bool, *models.HTTPError) {
	args := m.Called(id)

	return args[0].(bool), args[1].(*models.HTTPError)
}

// NewCompanyRepoMock generate new company repo mock
func NewCompanyRepoMock() *CompanyRepoMock {
	return &CompanyRepoMock{}
}
