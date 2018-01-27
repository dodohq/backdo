package models

// User accounts for Do-do's partners' staffs
type User struct {
	Email     string   `json:"email" schema:"email"`
	Password  string   `json:"password" schema:"password"`
	CompanyID int64    `json:"company_id" schema:"company_id"`
	Company   *Company `json:"company,omitempty"`
}
