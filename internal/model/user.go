package model

import "time"

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password_hashed"`
	Role      int       `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UpdateUserModel struct {
	ID       int64   `db:"id"`
	Username *string `db:"username"`
	Password *string `db:"password"`
}
