package user

import (
	"fmt"

	"github.com/dodohq/backdo/models"
)

// FetchAllUsers get all existing users
func (r *privateUserRepo) FetchAllUsers() ([]*models.User, *models.HTTPError) {
	query := `SELECT id, email, password, company_id FROM users WHERE NOT deleted`
	list, err := r.fetch(query)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	return list, nil
}

// InsertNewUser insert new user entry
func (r *privateUserRepo) InsertNewUser(u *models.User) (*models.User, *models.HTTPError) {
	query := `INSERT INTO users(email, password, company_id) VALUES($1, $2, $3) RETURNING id`
	var ID int64
	err := r.Conn.QueryRow(query, u.Email, u.Password, u.CompanyID).Scan(&ID)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	u.ID = ID
	return u, nil
}

// DeleteUser delete user
func (r *privateUserRepo) DeleteUser(id int64) *models.HTTPError {
	query := `UPDATE users SET deleted = TRUE WHERE id = $1 AND NOT deleted`
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
	}
	if rowsAffected < 1 {
		return models.NewErrorNotFound(fmt.Sprintf("User with ID %d doesn't exist", id))
	}

	return nil
}

func (r *privateUserRepo) fetch(query string, args ...interface{}) ([]*models.User, error) {
	rows, err := r.Conn.Query(query, args...)

	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]*models.User, 0)
	for rows.Next() {
		t := new(models.User)
		err = rows.Scan(&t.Email, &t.Password, &t.CompanyID)
		if err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}
