package repository

const (
	queryCreateAdmin = `
insert into auth."user" 
    (username, password_hashed, created_at, updated_at, role)
values 
	($1, $2, $3, $4, $5)
`
	queryGetUserByID = `
select
	username,
	password_hashed,
	created_at,
	updated_at,
	role
from auth."user"
where id = $1
`
	queryDeleteUser = `
delete
from auth."user"
where id = $1
`
	queryIsUserExists = `
select exists (
	select 1
	from auth."user"
	where username = $1
)
`
)
