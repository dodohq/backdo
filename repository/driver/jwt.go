package driver

import (
	"github.com/dodohq/backdo/lib/jwt"
	"github.com/dodohq/backdo/models"
)

// GenerateJWT generate jwt from driver object
func (r *privateDriverRepo) GenerateJWT(d *models.Driver) (string, *models.HTTPError) {
	token, err := jwt.CreateToken(*d, jwt.DriverType)
	if err != nil {
		return "", models.NewErrorInternalServer(err.Error())
	}

	return token, nil
}
