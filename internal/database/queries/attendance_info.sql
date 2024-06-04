-- name: InsertAttendance :exec
INSERT INTO attendance_info (student_id, date, status, class_id, division_id)
VALUES ($1, $2, $3, $4, $5);

-- name: GetAttendanceByStudent :many
SELECT * FROM attendance_info
WHERE student_id = $1
ORDER BY date DESC;

-- name: GetAttendanceByClassDivisionAndMonthYear :many
SELECT * FROM attendance_info
WHERE class_id = $1 AND division_id = $2 AND DATE_TRUNC('month', date) = $3
ORDER BY date DESC;

