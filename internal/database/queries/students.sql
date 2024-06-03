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

-- name: UpdateStudentInfo :one
UPDATE students
SET firstName = $2,
    lastName = $3,
    rollno = $4,
    email = $5,
    className = $6,
    division = $7,
    year = $8
WHERE student_id = $1
RETURNING student_id;

-- name: DeleteStudentInfo :one
DELETE FROM students
WHERE student_id = $1
RETURNING student_id;