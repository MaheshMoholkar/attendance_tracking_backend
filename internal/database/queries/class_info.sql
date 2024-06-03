-- name: CreateClassInfo :one
INSERT INTO class_info (className) 
VALUES ($1) 
RETURNING class_id;

-- name: GetClasses :many
SELECT * 
FROM class_info;

-- name: UpdateClassInfo :one
UPDATE class_info 
SET className = $1 
WHERE class_id = $2 
RETURNING class_id;

-- name: DeleteClassInfo :exec
DELETE FROM class_info 
WHERE class_id = $1;

-- name: GetClassDivisions :many
SELECT c.className, d.divisionName 
FROM class_info c
JOIN division_info d ON c.class_id = d.class_id;
