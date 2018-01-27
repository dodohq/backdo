package models

// Company model to represent Do-do's logistics partner
type Company struct {
	ID            int64  `json:"id"`
	Name          string `json:"name" schema:"name"`
	ContactNumber string `json:"contact_number" schema:"contact_number"`
}
