-- name: CreateStudentCredentials :one
INSERT INTO student_credentials(
    student_id,
    password_hash
)
VALUES ($1, $2)
RETURNING student_id;

-- name: UpdateStudentCredentials :one
UPDATE student_credentials
SET 
    student_id = $2,
    password_hash = $3
WHERE student_id = $1
RETURNING student_id;

-- name: GetStudentCredentials :one
SELECT password_hash
FROM student_credentials
WHERE student_id = $1;