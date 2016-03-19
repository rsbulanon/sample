package models

type User struct {
	Model
	FirstName  string `json:"first_name" binding:"required"`
	LastName string `json:"last_name,omitempty"`
	ContactNo string `json:"contact_no,omitempty"`
	Email    string `json:"email,omitempty"`
	Status string `json:"status,omitempty" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}