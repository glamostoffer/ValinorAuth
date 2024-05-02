package repository

const (
	queryCreateUser = `
insert into auth."user" 
    (username, password_hashed, created_at, updated_at, role)
values 
	($1, $2, $3, $4, 1)
`
)
