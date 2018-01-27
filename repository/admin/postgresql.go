package admin

import (
	"fmt"

	"github.com/dodohq/backdo/models"
)

func (r *privateAdminRepo) fetch(query string, args ...interface{}) ([]*models.Admin, *models.HTTPError) {
	rows, err := r.Conn.Query(query, args...)

	if err != nil {
		fmt.Println("ERROR DB:", err)
		return nil, models.NewErrorInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]*models.Admin, 0)
	for rows.Next() {
		t := new(models.Admin)
		err = rows.Scan(&t.Email, &t.Password)

		if err != nil {
			fmt.Println("ERROR DB:", err)
			return nil, models.NewErrorInternalServer(err.Error())
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

	a := &models.Admin{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NewErrorNotFound("Email Not Found")
	}

	return a, nil
}
