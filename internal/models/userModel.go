package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}
