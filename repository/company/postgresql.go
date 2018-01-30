package company

import (
	"fmt"

	"github.com/dodohq/backdo/models"
)

func (r *privateCompanyRepo) fetch(query string, args ...interface{}) ([]*models.Company, error) {
	rows, err := r.Conn.Query(query, args...)

	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]*models.Company, 0)
	for rows.Next() {
		t := new(models.Company)
		err := rows.Scan(&t.ID, &t.Name, &t.ContactNumber)

		if err != nil {
			return nil, models.NewErrorInternalServer(err.Error())
		}

		results = append(results, t)
	}

	return results, nil
}

// GetAllCompany get all partner companies
func (r *privateCompanyRepo) GetAllCompany() ([]*models.Company, *models.HTTPError) {
	query := `SELECT id, name, contact_number FROM companies WHERE NOT deleted`
	companiesList, err := r.fetch(query)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	return companiesList, nil
}

// InsertNewCompany onboard a new partner company
func (r *privateCompanyRepo) InsertNewCompany(c *models.Company) (*models.Company, *models.HTTPError) {
	query := `INSERT INTO companies (name, contact_number) VALUES ($1, $2) RETURNING id`
	var ID int64
	err := r.Conn.QueryRow(query, c.Name, c.ContactNumber).Scan(&ID)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	c.ID = ID
	return c, nil
}

// DeleteACompany delete a company from db
func (r *privateCompanyRepo) DeleteACompany(id int64) (bool, *models.HTTPError) {
	query := `UPDATE companies SET deleted = TRUE WHERE id = $1 AND NOT deleted`
	stmt, err := r.Conn.Prepare(query)
	if err != nil {
		return false, models.NewErrorInternalServer(err.Error())
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, models.NewErrorInternalServer(err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, models.NewErrorInternalServer(err.Error())
	}
	if rowsAffected < 1 {
		return false, models.NewErrorNotFound(fmt.Sprintf("Company with ID %d doesn't exist", id))
	}

	return true, nil
}
