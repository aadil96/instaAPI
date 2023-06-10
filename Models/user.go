package models

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Mobile uint64 `json:"mobile"`
	Email string `json:"email"`
}