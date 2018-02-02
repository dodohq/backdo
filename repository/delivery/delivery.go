package delivery

import (
	"database/sql"

	"github.com/dodohq/backdo/repository"
)

type privateDeliveryRepo struct {
	Conn *sql.DB
}

// NewDeliveryRepository generate new delivery repository
func NewDeliveryRepository(Conn *sql.DB) repository.DeliveryRepository {
	return &privateDeliveryRepo{Conn}
}
