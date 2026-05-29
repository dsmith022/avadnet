-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    hashed_password
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET
    username = $2,
    email = $3,
    hashed_password = $4,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;