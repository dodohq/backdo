package models

// Delivery representation of each delivery needs to be made
type Delivery struct {
	ID            int64    `json:"id"`
	CustomerName  string   `json:"customer_name" schema:"customer_name"`
	ContactNumber string   `json:"contact_number" schema:"contact_number"`
	PassCode      string   `json:"passcode"`
	QRCodeURL     string   `json:"qr_code_url"`
	CompanyID     int64    `json:"company_id" schema:"company_id"`
	Company       *Company `json:"company,omitempty"`
}
