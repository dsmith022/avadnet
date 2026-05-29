-- name: CreateVenture :one
INSERT INTO ventures (
    name,
    description,
    created_by_user_id
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetVentureByID :one
SELECT *
FROM ventures
WHERE id = $1;

-- name: ListVentures :many
SELECT *
FROM ventures
ORDER BY created_at DESC;

-- name: AddVentureMember :one
INSERT INTO venture_members (
    venture_id,
    user_id,
    role
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListVentureMembers :many
SELECT *
FROM venture_members
WHERE venture_id = $1
ORDER BY created_at ASC;