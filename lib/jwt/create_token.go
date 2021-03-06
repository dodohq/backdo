package jwt

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dodohq/backdo/models"
)

// CreateToken create jwt for either admin or user
func CreateToken(v interface{}, vType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix() // 1 week

	if vType == AdminType {
		claims["id"] = v.(models.Admin).ID
		claims["email"] = v.(models.Admin).Email
		claims["is_admin"] = true
	} else if vType == UserType {
		claims["email"] = v.(models.User).Email
		claims["company_id"] = v.(models.User).CompanyID
		claims["is_admin"] = false
	} else {
		return "", errors.New("Interface Type Not Supported")
	}

	token.Claims = claims
	return token.SignedString([]byte(tokenEncodeString))
}
