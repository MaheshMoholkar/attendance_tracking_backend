-- name: CreateStudentInfo :one
INSERT INTO student_info (
    firstName,
    lastName,
    rollno,
    email,
    class_id,
    division_id,
    academic_year,
    year,
    student_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING student_id;

-- name: GetStudentInfo :one
SELECT * 
FROM student_info 
WHERE student_id = $1;

-- name: GetStudentsInfo :many
SELECT * 
FROM student_info;

-- name: UpdateStudentInfo :exec
UPDATE student_info
SET firstName = $2,
    lastName = $3,
    rollno = $4,
    email = $5,
    class_id = $6,
    division_id = $7,
    academic_year = $8,
    year = $9
WHERE student_id = $1;

-- name: DeleteStudentInfo :exec
DELETE FROM student_info
WHERE student_id = $1;