package repomock

import (
	"github.com/dodohq/backdo/models"
	"github.com/stretchr/testify/mock"
)

// AdminRepoMock mock of admin repo
type AdminRepoMock struct {
	mock.Mock
	UniversalPassword string
	NonExistentEmail  string
}

// GetByEmail mock of func
func (m *AdminRepoMock) GetByEmail(email string) (*models.Admin, *models.HTTPError) {
	args := m.Called(email)
	if args.String(0) == m.NonExistentEmail {
		return nil, models.NewErrorNotFound("Email Not Found")
	}

	return &models.Admin{Email: args.String(0), Password: m.UniversalPassword}, nil
}

// NewAdminRepoMock generate new mock
func NewAdminRepoMock(up, nem string) *AdminRepoMock {
	return &AdminRepoMock{
		UniversalPassword: up,
		NonExistentEmail:  nem,
	}
}
