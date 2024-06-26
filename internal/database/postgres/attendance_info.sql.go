// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: attendance_info.sql

package postgres

import (
	"context"
)

const fetchAttendanceTableName = `-- name: FetchAttendanceTableName :one
SELECT attendance_table_name
FROM attendance_info
WHERE class_id = $1 AND division_id = $2 AND attendance_month_year = $3
`

type FetchAttendanceTableNameParams struct {
	ClassID             int32
	DivisionID          int32
	AttendanceMonthYear string
}

func (q *Queries) FetchAttendanceTableName(ctx context.Context, arg FetchAttendanceTableNameParams) (string, error) {
	row := q.db.QueryRowContext(ctx, fetchAttendanceTableName, arg.ClassID, arg.DivisionID, arg.AttendanceMonthYear)
	var attendance_table_name string
	err := row.Scan(&attendance_table_name)
	return attendance_table_name, err
}
