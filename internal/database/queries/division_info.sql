-- name: GetDivisions :many
SELECT * 
FROM division_info;

-- name: CreateDivisionInfo :one
INSERT INTO division_info (division_name, class_id) 
VALUES ($1, $2) 
RETURNING *;

-- name: UpdateDivisionInfo :one
UPDATE division_info 
SET division_name = $1, class_id = $2 
WHERE division_id = $3 
RETURNING *;

-- name: DeleteDivisionInfo :exec
DELETE FROM division_info 
WHERE division_id = $1;