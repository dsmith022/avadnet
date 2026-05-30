-- name: GetChurchById :one
SELECT * FROM churches WHERE id = $1;

-- name: ListChurches :many
SELECT * FROM churches ORDER BY name;

-- name: CreateChurch :one
INSERT INTO churches (name, description) VALUES ($1, $2) RETURNING *;

-- name: UpdateChurch :one
UPDATE churches SET name = $2, description = $3, updated_at = now() WHERE id = $1 RETURNING *;

-- name: DeleteChurch :exec
DELETE FROM churches WHERE id = $1;

-- name: AddChurchMember :one
INSERT INTO church_members (church_id, user_id, role) VALUES ($1, $2, $3) RETURNING *;

-- name: ListChurchMembers :many
SELECT
    cm.*,
    u.username,
    u.email,
    u.id AS user_id
FROM church_members cm
JOIN users u ON u.id = cm.user_id
WHERE cm.church_id = $1
ORDER BY cm.created_at ASC;     

-- name: ListUserChurches :many
SELECT
    c.name AS church_name,
    c.id AS church_id,
    cm.role AS member_role
FROM church_members cm
JOIN churches c ON c.id = cm.church_id
WHERE cm.user_id = $1
ORDER BY cm.created_at ASC;