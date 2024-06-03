-- name: CreateStaffInfo :one
INSERT INTO staff_info (
    firstName,
    lastName,
    email,
    staff_id
)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetStaffInfo :one
SELECT * 
FROM staff_info 
WHERE staff_id = $1;

-- name: GetStaffsInfo :many
SELECT * 
FROM staff_info;

-- name: UpdateStaffInfo :one
UPDATE staff_info
SET firstName = $2,
    lastName = $3,
    email = $4
WHERE staff_id = $1
RETURNING id;

-- name: DeleteStaffInfo :one
DELETE FROM staff_info
WHERE staff_id = $1
RETURNING id;
