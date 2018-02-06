package models

// Driver accounts for Do-do's partners' drivers
type Driver struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name" schema:"name"`
	PhoneNumber string   `json:"phone_number" schema:"phone_number"`
	CompanyID   int64    `json:"company_id"`
	Company     *Company `json:"company,omitempty"`
}
