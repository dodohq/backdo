package usecase_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	testTable := []struct {
		Email    string
		Password string
	}{
		{"stanley@dodo.tech", adminRepoMock.UniversalPassword},
		{"hello@dodo.tech", "notuniversalpassword"},
		{adminRepoMock.NonExistentEmail, adminRepoMock.UniversalPassword},
	}

	for i, fakeAccDetails := range testTable {
		adminRepoMock.On("GetByEmail", fakeAccDetails.Email).Return(fakeAccDetails.Email)
		adminRepoMock.On("GenerateJWT").Return()
		actual, err := adminUsecase.Login(fakeAccDetails.Email, fakeAccDetails.Password)

		if i == 0 {
			if err != nil {
				t.Errorf(err.Error())
			} else {
				assert.Equal(t, actual, adminRepoMock.FakeJWT)
			}
		} else if i == 1 && err.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected %d, Got %d. Error Msg: %s", http.StatusUnauthorized, err.StatusCode, err.Error())
		} else if i == 2 && err.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %d, Got %d. Error Msg: %s", http.StatusNotFound, err.StatusCode, err.Error())
		}
	}
}
