-- name: CreateStaffCredentials :one
INSERT INTO staff_credentials(
    staff_id,
    password_hash
)
VALUES ($1, $2)
RETURNING staff_id;

-- name: UpdateStaffCredentials :one
UPDATE staff_credentials
SET 
    staff_id = $2,
    password_hash = $3
WHERE staff_id = $1
RETURNING staff_id;

-- name: GetStaffCredentials :one
SELECT password_hash 
FROM staff_credentials
WHERE staff_id = $1;