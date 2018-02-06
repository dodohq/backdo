package driver

import (
	"fmt"

	"github.com/dodohq/backdo/models"
)

// FetchDriversByCompany get all drivers registered under a company
func (r *privateDriverRepo) FetchDriversByCompany(companyID int64) ([]*models.Driver, *models.HTTPError) {
	query := `SELECT id, name, phone_number, company_id FROM drivers WHERE company_id = $1 AND NOT deleted`
	list, err := r.fetch(query, companyID)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	return list, nil
}

// FetchDriverByID get a driver info using ID
func (r *privateDriverRepo) FetchDriverByID(id int64) (*models.Driver, *models.HTTPError) {
	query := `SELECT id, name, phone_number, company_id FROM drivers WHERE id = $1 AND NOT deleted`
	list, err := r.fetch(query, id)
	if err != nil {
		return nil, models.NewErrorUnprocessableEntity(err.Error())
	} else if len(list) < 1 {
		return nil, nil
	}

	return list[0], nil
}

// FetchDriverByPhoneNumber get a driver info using phone number
func (r *privateDriverRepo) FetchDriverByPhoneNumber(phoneNo string) (*models.Driver, *models.HTTPError) {
	query := `SELECT id, name, phone_number, company_id FROM drivers WHERE phone_number = $1 AND NOT deleted`
	list, err := r.fetch(query, phoneNo)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	} else if len(list) < 1 {
		return nil, nil
	}

	return list[0], nil
}

// InsertNewDriver insert new driver entry to db
func (r *privateDriverRepo) InsertNewDriver(d *models.Driver) (*models.Driver, *models.HTTPError) {
	query := `INSERT INTO drivers(name, phone_number, company_id) VALUES($1, $2, $3) RETURNING id`
	var ID int64
	err := r.Conn.QueryRow(query, d.Name, d.PhoneNumber, d.CompanyID).Scan(&ID)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	d.ID = ID
	return d, nil
}

// DeleteDriver delete driver
func (r *privateDriverRepo) DeleteDriver(id int64) *models.HTTPError {
	query := `UPDATE drivers SET deleted = TRUE WHERE id = $1 AND NOT deleted`
	stmt, err := r.Conn.Prepare(query)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	} else if rowsAffected < 1 {
		return models.NewErrorNotFound(fmt.Sprintf("Driver with id %d doesn't exist", id))
	}

	return nil
}

func (r *privateDriverRepo) fetch(query string, args ...interface{}) ([]*models.Driver, error) {
	rows, err := r.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*models.Driver, 0)
	for rows.Next() {
		d := new(models.Driver)
		err = rows.Scan(&d.ID, &d.Name, &d.PhoneNumber, &d.CompanyID)
		if err != nil {
			return nil, err
		}
		results = append(results, d)
	}

	return results, nil
}
