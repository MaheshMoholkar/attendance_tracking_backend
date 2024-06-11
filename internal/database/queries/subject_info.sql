-- name: GetSubjects :many
SELECT * 
FROM subject_info;

-- name: CreateSubjectInfo :one
INSERT INTO subject_info (subjectName, class_id) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetSubjectIDByName :one
SELECT subject_id 
FROM subject_info
WHERE subjectName = $1 AND class_id=$2;

-- name: UpdateSubjectInfo :one
UPDATE subject_info 
SET subjectName = $1, class_id = $2 
WHERE subject_id = $3 
RETURNING *;

-- name: DeleteSubjectInfo :exec
DELETE FROM subject_info 
WHERE subject_id = $1;
