package admin

import (
	"github.com/dodohq/backdo/models"
)

func (r *privateAdminRepo) fetch(query string, args ...interface{}) ([]*models.Admin, error) {
	rows, err := r.Conn.Query(query, args...)

	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]*models.Admin, 0)
	for rows.Next() {
		t := new(models.Admin)
		err = rows.Scan(&t.ID, &t.Email, &t.Password)

		if err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, nil
}

// GetByEmail get admin account by admin
func (r *privateAdminRepo) GetByEmail(email string) (*models.Admin, *models.HTTPError) {
	query := `SELECT * FROM admins WHERE email = $1`
	list, err := r.fetch(query, email)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	if len(list) < 1 {
		return nil, nil
	}

	return list[0], nil
}
