// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: student_credentials.sql

package postgres

import (
	"context"
)

const createStudentCredentials = `-- name: CreateStudentCredentials :one
INSERT INTO student_credentials(
    student_id,
    password_hash
)
VALUES ($1, $2)
RETURNING student_id
`

type CreateStudentCredentialsParams struct {
	StudentID    int32
	PasswordHash string
}

func (q *Queries) CreateStudentCredentials(ctx context.Context, arg CreateStudentCredentialsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createStudentCredentials, arg.StudentID, arg.PasswordHash)
	var student_id int32
	err := row.Scan(&student_id)
	return student_id, err
}

const getStudentCredentials = `-- name: GetStudentCredentials :one
SELECT password_hash
FROM student_credentials
WHERE student_id = $1
`

func (q *Queries) GetStudentCredentials(ctx context.Context, studentID int32) (string, error) {
	row := q.db.QueryRowContext(ctx, getStudentCredentials, studentID)
	var password_hash string
	err := row.Scan(&password_hash)
	return password_hash, err
}

const updateStudentCredentials = `-- name: UpdateStudentCredentials :one
UPDATE student_credentials
SET 
    student_id = $2,
    password_hash = $3
WHERE student_id = $1
RETURNING student_id
`

type UpdateStudentCredentialsParams struct {
	StudentID    int32
	StudentID_2  int32
	PasswordHash string
}

func (q *Queries) UpdateStudentCredentials(ctx context.Context, arg UpdateStudentCredentialsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateStudentCredentials, arg.StudentID, arg.StudentID_2, arg.PasswordHash)
	var student_id int32
	err := row.Scan(&student_id)
	return student_id, err
}
