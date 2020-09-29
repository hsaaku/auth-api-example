package model

// User ...
type User struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	DisplayName string `json:"display_name" db:"display_name"`
	Password    string `json:"password,omitempty" db:"password"`
}
