-- name: GetDivisions :many
SELECT * 
FROM division_info;

-- name: CreateDivisionInfo :one
INSERT INTO division_info (divisionName, class_id) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetDivisionIDByName :one
SELECT division_id 
FROM division_info
WHERE divisionName = $1 AND class_id=$2;

-- name: UpdateDivisionInfo :one
UPDATE division_info 
SET divisionName = $1, class_id = $2 
WHERE division_id = $3 
RETURNING *;

-- name: DeleteDivisionInfo :exec
DELETE FROM division_info 
WHERE division_id = $1;