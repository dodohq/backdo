package repomock

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/dodohq/backdo/models"
	"github.com/stretchr/testify/mock"
)

// AdminRepoMock mock of admin repo
type AdminRepoMock struct {
	mock.Mock
	UniversalPassword string
	NonExistentEmail  string
	FakeJWT           string
}

// GetByEmail mock of func
func (m *AdminRepoMock) GetByEmail(email string) (*models.Admin, *models.HTTPError) {
	args := m.Called(email)
	if args.String(0) == m.NonExistentEmail {
		return nil, nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(m.UniversalPassword), 14)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	return &models.Admin{Email: args.String(0), Password: string(hash)}, nil
}

// GenerateJWT mock of func
func (m *AdminRepoMock) GenerateJWT(_ *models.Admin) (string, *models.HTTPError) {
	m.Called()

	return m.FakeJWT, nil
}

// NewAdminRepoMock generate new mock
func NewAdminRepoMock(up, nem, jwt string) *AdminRepoMock {
	return &AdminRepoMock{
		UniversalPassword: up,
		NonExistentEmail:  nem,
		FakeJWT:           jwt,
	}
}
