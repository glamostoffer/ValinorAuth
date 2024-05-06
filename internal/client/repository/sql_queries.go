package repository

const (
	queryCreateUser = `
insert into auth."user" 
    (username, password_hashed, created_at, updated_at, role)
values 
	($1, $2, $3, $4, $5)
`
	queryGetUserByID = `
select
username, password_hashed, created_at, updated_at, role
from auth."user"
where id = $1
`
	queryGetUserByLogin = `
select
id, username, password_hashed, created_at, updated_at, role
from auth."user"
where username = $1
`
	queryIsUserExists = `
select exists (
	select 1
	from auth."user"
	where username = $1
)
`
	queryUpdateUser = `
update auth."user" 
set 
    username = coalesce($2, username),
 	password_hashed = coalesce($3, password_hashed),
 	updated_at = $4
where
    id = $1
`
)
