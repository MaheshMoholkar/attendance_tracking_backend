// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: students.sql

package postgres

import (
	"context"
)

const createStudentInfo = `-- name: CreateStudentInfo :one
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
RETURNING student_id
`

type CreateStudentInfoParams struct {
	Firstname string
	Lastname  string
	Rollno    int32
	Email     string
	Classname string
	Division  string
	Year      int32
	StudentID int32
}

func (q *Queries) CreateStudentInfo(ctx context.Context, arg CreateStudentInfoParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createStudentInfo,
		arg.Firstname,
		arg.Lastname,
		arg.Rollno,
		arg.Email,
		arg.Classname,
		arg.Division,
		arg.Year,
		arg.StudentID,
	)
	var student_id int32
	err := row.Scan(&student_id)
	return student_id, err
}

const getStudentInfo = `-- name: GetStudentInfo :one
SELECT id, firstname, lastname, rollno, email, classname, division, year, student_id 
FROM students 
WHERE student_id = $1
`

func (q *Queries) GetStudentInfo(ctx context.Context, studentID int32) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentInfo, studentID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Rollno,
		&i.Email,
		&i.Classname,
		&i.Division,
		&i.Year,
		&i.StudentID,
	)
	return i, err
}

const getStudentsInfo = `-- name: GetStudentsInfo :many
SELECT id, firstname, lastname, rollno, email, classname, division, year, student_id 
FROM students
`

func (q *Queries) GetStudentsInfo(ctx context.Context) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getStudentsInfo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Lastname,
			&i.Rollno,
			&i.Email,
			&i.Classname,
			&i.Division,
			&i.Year,
			&i.StudentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
