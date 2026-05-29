-- name: CreateMissionOrganization :one
INSERT INTO mission_organizations (
    name,
    description
)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetMissionOrganizationByID :one
SELECT *
FROM mission_organizations
WHERE id = $1;

-- name: ListMissionOrganizations :many
SELECT *
FROM mission_organizations
ORDER BY created_at DESC;

-- name: AddMissionMember :one
INSERT INTO mission_members (
    mission_id,
    user_id,
    role
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListMissionMembers :many
SELECT *
FROM mission_members
WHERE mission_id = $1
ORDER BY created_at ASC;

-- name: AddMissionVenture :one
INSERT INTO mission_ventures (
    mission_id,
    venture_id,
    description
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListMissionVentures :many
SELECT
    mv.*,
    v.name AS venture_name
FROM mission_ventures mv
JOIN ventures v ON v.id = mv.venture_id
WHERE mv.mission_id = $1
ORDER BY mv.created_at ASC;