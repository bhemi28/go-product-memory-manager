-- name: CreateUser :one
INSERT INTO users (id, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByMail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetByUserName :one
SELECT * FROM users
WHERE username = $1;
