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

// GetUserByEmail get user by email
func (r *privateUserRepo) GetUserByEmail(email string) (*models.User, *models.HTTPError) {
	query := `SELECT id, email, password, company_id FROM users WHERE email = $1 AND NOT deleted`
	list, err := r.fetch(query, email)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	if len(list) < 1 {
		return nil, nil
	}

	u := list[0]
	return u, nil
}

// UpdateUser update user profile info
func (r *privateUserRepo) UpdateUser(u *models.User) *models.HTTPError {
	query := `UPDATE users SET email = $1, password = $2, company_id = $3 WHERE id = $4 AND NOT deleted`
	stmt, err := r.Conn.Prepare(query)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}

	res, err := stmt.Exec(u.Email, u.Password, u.CompanyID, u.ID)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	} else if rowsAffected < 1 {
		return models.NewErrorNotFound("No Existing User")
	}

	return nil
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
		return nil, err
	}
	defer rows.Close()

	results := make([]*models.User, 0)
	for rows.Next() {
		t := new(models.User)
		err = rows.Scan(&t.ID, &t.Email, &t.Password, &t.CompanyID)
		if err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}
