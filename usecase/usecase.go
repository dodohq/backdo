package usecase

import "github.com/dodohq/backdo/models"

// AdminUsecase usecase interface for handling admin endpoints
type AdminUsecase interface {
	Login(email, password string) (string, *models.HTTPError)
}
