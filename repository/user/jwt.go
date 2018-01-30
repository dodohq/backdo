package user

import (
	"github.com/dodohq/backdo/lib/jwt"
	"github.com/dodohq/backdo/models"
)

// GenerateJWT generate jwt from user object
func (r *privateUserRepo) GenerateJWT(u *models.User) (string, *models.HTTPError) {
	token, err := jwt.CreateToken(*u, jwt.UserType)
	if err != nil {
		return "", models.NewErrorInternalServer(err.Error())
	}

	return token, nil
}
