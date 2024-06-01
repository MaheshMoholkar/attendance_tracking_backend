-- name: CreateStudentInfo :one
INSERT INTO students (
    firstName,
    lastName,
    rollno,
    email,
    className,
    division,
    year,
    student_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING student_id;

-- name: GetStudentInfo :one
SELECT * 
FROM students 
WHERE student_id = $1;

-- name: GetStudentsInfo :many
SELECT * 
FROM students;