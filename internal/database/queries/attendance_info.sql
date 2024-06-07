-- name: FetchAttendanceTableName :one
SELECT attendance_table_name
FROM attendance_info
WHERE class_id = $1 AND division_id = $2 AND attendance_month_year = $3;
