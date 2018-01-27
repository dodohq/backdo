package usecase_test

import (
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
		actual, err := adminUsecase.Login(fakeAccDetails.Email, fakeAccDetails.Password)
		if i == 0 {
			if err != nil {
				t.Errorf(err.Error())
			} else {
				assert.Equal(t, actual, fakeAccDetails)
			}
		} else if i == 1 && err.StatusCode != 403 {
			t.Errorf("Expected 403, Got %d. Error Msg: %s", err.StatusCode, err.Error())
		} else if i == 2 && err.StatusCode != 404 {
			t.Errorf("Expected 404, Got %d. Error Msg: %s", err.StatusCode, err.Error())
		}
	}
}
