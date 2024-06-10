package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	db    *sql.DB
	store *database.Store
}

func NewAttendanceHandler(store *database.Store) *AttendanceHandler {
	dbConn := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}
	return &AttendanceHandler{
		db:    conn,
		store: store,
	}
}

func (h *AttendanceHandler) InitializeAttendanceTableHandler(ctx *fiber.Ctx) error {
	className := ctx.Query("class_name")
	divisionName := ctx.Query("division_name")
	subject := ctx.Query("subject")
	monthYear := ctx.Query("month_year")

	classID, err := h.store.DB.GetClassIDByName(ctx.Context(), className)
	if err != nil {
		return err
	}
	divisionID, err := h.store.DB.GetDivisionIDByName(ctx.Context(), postgres.GetDivisionIDByNameParams{
		Divisionname: divisionName,
		ClassID:      classID,
	})
	if err != nil {
		return err
	}

	tableName := fmt.Sprintf("%s_%d_%d_%s", subject, classID, divisionID, monthYear)

	// SQL statement to create the attendance table with a unique constraint
	createTableStatement := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		student_id INT NOT NULL,
		attendance_date DATE NOT NULL,
		attendance_status TEXT NOT NULL,
		UNIQUE (student_id, attendance_date)
	)
	`, tableName)

	_, err = h.db.Exec(createTableStatement)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *AttendanceHandler) UpdateAttendanceHandler(ctx *fiber.Ctx) error {
	var update types.AttendanceUpdate
	if err := ctx.BodyParser(&update); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	className := update.ClassName
	divisionName := update.DivisionName
	subject := update.Subject
	monthYear := update.MonthYear

	classID, err := h.store.DB.GetClassIDByName(ctx.Context(), className)
	if err != nil {
		return err
	}
	divisionID, err := h.store.DB.GetDivisionIDByName(ctx.Context(), postgres.GetDivisionIDByNameParams{
		Divisionname: divisionName,
		ClassID:      classID,
	})
	if err != nil {
		return err
	}

	tableName := fmt.Sprintf("%s_%d_%d_%s", subject, classID, divisionID, monthYear)
	if len(tableName) == 0 {
		return ctx.SendString("invalid table_name")
	}

	log.Print(tableName)

	// Insert attendance data
	for _, studentAttendance := range update.RowData {
		studentID := studentAttendance.StudentID
		for day, status := range studentAttendance.Attendance {
			attendanceDate := time.Date(time.Now().Year(), time.Now().Month(), day, 0, 0, 0, 0, time.UTC)
			attendanceStatus := "absent"
			if status {
				attendanceStatus = "present"
			}

			insertQuery := fmt.Sprintf(`
                INSERT INTO %s (student_id, attendance_date, attendance_status)
                VALUES ($1, $2, $3)
                ON CONFLICT (student_id, attendance_date) DO UPDATE SET attendance_status = EXCLUDED.attendance_status
            `, tableName)

			if _, err := h.db.Exec(insertQuery, studentID, attendanceDate, attendanceStatus); err != nil {
				return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
			}
		}
	}

	return ctx.SendString("Attendance updated successfully")
}

func (h *AttendanceHandler) GetAttendanceHandler(ctx *fiber.Ctx) error {
	className := ctx.Query("class_name")
	divisionName := ctx.Query("division_name")
	subject := ctx.Query("subject")
	monthYear := ctx.Query("month_year")

	classID, err := h.store.DB.GetClassIDByName(ctx.Context(), className)
	if err != nil {
		return err
	}
	divisionID, err := h.store.DB.GetDivisionIDByName(ctx.Context(), postgres.GetDivisionIDByNameParams{
		Divisionname: divisionName,
		ClassID:      classID,
	})
	if err != nil {
		return err
	}

	tableName := fmt.Sprintf("%s_%d_%d_%s", subject, classID, divisionID, monthYear)

	log.Print(tableName)
	// SQL query to fetch attendance data
	query := fmt.Sprintf(`
        SELECT student_id, attendance_date, attendance_status
        FROM %s
    `, tableName)

	rows, err := h.db.Query(query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	defer rows.Close()

	// Process the rows and organize data into the required format
	attendanceData := make(map[int]map[int]bool)
	for rows.Next() {
		var studentID int
		var attendanceDate time.Time
		var attendanceStatus string
		if err := rows.Scan(&studentID, &attendanceDate, &attendanceStatus); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		}

		if _, exists := attendanceData[studentID]; !exists {
			attendanceData[studentID] = make(map[int]bool)
		}

		day := attendanceDate.Day()
		attendanceData[studentID][day] = attendanceStatus == "present"
	}

	if err := rows.Err(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// Format the attendance data into the required JSON structure
	var result []map[string]interface{}
	for studentID, attendance := range attendanceData {
		studentAttendance := map[string]interface{}{
			"student_id": studentID,
			"attendance": attendance,
		}
		result = append(result, studentAttendance)
	}

	return ctx.JSON(fiber.Map{
		"data": result,
	})
}
