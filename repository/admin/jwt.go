package admin

import (
	"github.com/dodohq/backdo/lib/jwt"
	"github.com/dodohq/backdo/models"
)

// GenerateJWT generate jwt from admin object
func (r *privateAdminRepo) GenerateJWT(a *models.Admin) (string, *models.HTTPError) {
	token, err := jwt.CreateToken(*a, true, false)
	if err != nil {
		return "", models.NewErrorInternalServer(err.Error())
	}

	return token, nil
}
