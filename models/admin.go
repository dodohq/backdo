package models

// Admin model to represent Do-do's admin accounts
type Admin struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" schema:"email"`
	Password string `json:"password" schema:"password"`
}
