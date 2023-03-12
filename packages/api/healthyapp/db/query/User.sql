-- name: LoginAccount :one
SELECT u."ID", u.first_name, u.last_name, u.email, u.password 
FROM public."User" as u
Where u.email = $1 AND u.password = $2;
