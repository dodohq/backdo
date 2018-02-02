package delivery

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/dodohq/backdo/models"
)

// GetAllDeliveries fetch all deliveries in db
func (r *privateDeliveryRepo) GetAllDeliveries() ([]*models.Delivery, *models.HTTPError) {
	query := `
		SELECT id, customer_name, contact_number, passcode, qr_code_url, company_id
		FROM deliveries WHERE NOT deleted`
	deliveries, err := r.fetch(query)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	return deliveries, nil
}

// GetDeliveriesByCompanyID fetch all deliveries that belongs to the same company
func (r *privateDeliveryRepo) GetDeliveriesByCompanyID(id int64) ([]*models.Delivery, *models.HTTPError) {
	query := `SELECT id, customer_name, contact_number, passcode, qr_code_url, company_id FROM deliveries WHERE company_id = $1 AND NOT deleted`
	deliveries, err := r.fetch(query, id)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	return deliveries, nil
}

// GetDeliveryByID fetch delivery by its ID
func (r *privateDeliveryRepo) GetDeliveryByID(id int64) (*models.Delivery, *models.HTTPError) {
	query := `SELECT id, customer_name, contact_number, passcode, qr_code_url, company_id FROM deliveries WHERE id = $1 AND NOT deleted`
	list, err := r.fetch(query, id)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	} else if len(list) < 1 {
		return nil, nil
	}

	return list[0], nil
}

// InsertDelivery insert a new record
func (r *privateDeliveryRepo) InsertDelivery(d *models.Delivery) (*models.Delivery, *models.HTTPError) {
	query := `INSERT INTO deliveries(customer_name, contact_number, company_id) VALUES($1, $2, $3) RETURNING id`
	var ID int64
	err := r.Conn.QueryRow(query, d.CustomerName, d.ContactNumber, d.CompanyID).Scan(&ID)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}

	d.ID = ID
	return d, nil
}

// BulkInsertDelivery insert multiple records of deliveries
func (r *privateDeliveryRepo) BulkInsertDelivery(deliveryList []*models.Delivery) ([]*models.Delivery, *models.HTTPError) {
	query := `INSERT INTO deliveries(customer_name, contact_number, company_id) VALUES`
	queryParams := []interface{}{}
	for i := 0; i < len(deliveryList); i++ {
		query += "($" + strconv.Itoa(i*3+1) + ", $" + strconv.Itoa(i*3+2) + ", $" + strconv.Itoa(i*3+3) + ")"
		if i < len(deliveryList)-1 {
			query += ",\n"
		}
		queryParams = append(queryParams, deliveryList[i].CustomerName, deliveryList[i].ContactNumber, deliveryList[i].CompanyID)
	}
	query += " RETURNING id"

	rows, err := r.Conn.Query(query, queryParams...)
	if err != nil {
		return nil, models.NewErrorInternalServer(err.Error())
	}
	defer rows.Close()
	var idx int64
	for rows.Next() {
		err := rows.Scan(&deliveryList[idx].ID)
		if err != nil {
			return nil, models.NewErrorInternalServer(err.Error())
		}
		idx++
	}

	return deliveryList, nil
}

// DeleteDelivery delete a delivery record
func (r *privateDeliveryRepo) DeleteDelivery(id int64) *models.HTTPError {
	query := `UPDATE deliveries SET deleted = TRUE WHERE id = $1 AND NOT deleted`
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
		return models.NewErrorNotFound(fmt.Sprintf("Delivery with ID %d doesnt exist", id))
	}

	return nil
}

func (r *privateDeliveryRepo) fetch(query string, args ...interface{}) ([]*models.Delivery, error) {
	rows, err := r.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*models.Delivery, 0)
	for rows.Next() {
		d := new(models.Delivery)
		var passcode sql.NullString
		var qrcode sql.NullString
		err = rows.Scan(&d.ID, &d.CustomerName, &d.ContactNumber, &passcode, &qrcode, &d.CompanyID)
		if passcode.Valid {
			d.PassCode = passcode.String
		} else {
			d.PassCode = ""
		}
		if qrcode.Valid {
			d.QRCodeURL = qrcode.String
		} else {
			d.QRCodeURL = ""
		}
		if err != nil {
			return nil, err
		}
		results = append(results, d)
	}

	return results, nil
}
