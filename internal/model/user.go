package model

import "time"

type User struct {
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      int       `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserDetails struct {
	Name      string
	Surname   string
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      string
}
