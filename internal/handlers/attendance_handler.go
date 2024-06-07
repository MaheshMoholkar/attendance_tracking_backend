package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	store *database.Store
}

func NewAttendanceHandler(store *database.Store) *AttendanceHandler {
	return &AttendanceHandler{
		store: store,
	}
}

func (h *AttendanceHandler) InitializeAttendanceTableHandler(ctx *fiber.Ctx) error {
	classID := ctx.Params("class_id")
	divisionID := ctx.Params("division_id")
	monthYear := ctx.Params("month_year")

	dbURL := fmt.Sprint(os.Getenv("DB_URL"))
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlQuery := fmt.Sprintf(`
	DO $$
	DECLARE
		table_name TEXT := '%s' || '-' || '%s' || '-' || '%s';
		day_count INT;
		column_definitions TEXT;
		student RECORD;
	BEGIN
		day_count := EXTRACT(DAY FROM DATE_TRUNC('MONTH', '%s'::DATE) + INTERVAL '1 MONTH' - INTERVAL '1 DAY');
		SELECT string_agg('"' || generate_series(1, day_count)::TEXT || '" TEXT DEFAULT NULL', ', ') INTO column_definitions;
		IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = table_name) THEN
			EXECUTE 'CREATE TABLE ' || table_name || ' (student_id INT PRIMARY KEY, ' || column_definitions || ')';
			FOR student IN SELECT student_id FROM student_info WHERE class_id = %s AND division_id = %s LOOP
				EXECUTE 'INSERT INTO ' || table_name || ' (student_id) VALUES ($1)' USING student.student_id;
			END LOOP;
			INSERT INTO attendance_info (attendance_table_name, attendance_month_year, class_id, division_id) VALUES (table_name, '%s', %s, %s);
		END IF;
	END;
	$$ LANGUAGE plpgsql;
`, classID, divisionID, monthYear, monthYear, classID, divisionID, monthYear, classID, divisionID)

	// Execute the SQL query with the arguments
	_, err = db.ExecContext(ctx.Context(), sqlQuery, classID, divisionID, monthYear)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *AttendanceHandler) UpdateAttendanceHandler(c *fiber.Ctx) error {
	var update types.AttendanceUpdate
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	classID, err := c.ParamsInt("class_id")
	if err != nil {
		return err
	}
	divisionID, err := c.ParamsInt("division_id")
	if err != nil {
		return err
	}
	monthYear := c.Params("month_year")

	dbURL := fmt.Sprintf(":%s", os.Getenv("DB_URL"))
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tableName, err := h.fetchAttendanceTableName(context.Background(), int32(classID), int32(divisionID), monthYear)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not fetch attendance table name")
	}

	for day, status := range update.Attendance {
		dayCol := fmt.Sprint(day)
		sqlQuery := fmt.Sprintf(`
			UPDATE %s 
			SET %s = $1 
			WHERE student_id = $2;
			`, tableName, dayCol)

		// Execute the SQL query with the arguments
		_, err := db.ExecContext(c.Context(), sqlQuery, status, update.StudentID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not update attendance")
		}
	}

	return c.SendString("Attendance updated successfully")
}

func (h *AttendanceHandler) fetchAttendanceTableName(ctx context.Context, classID, divisionID int32, monthYear string) (string, error) {

	result, err := h.store.DB.FetchAttendanceTableName(ctx, postgres.FetchAttendanceTableNameParams{
		ClassID:             classID,
		DivisionID:          divisionID,
		AttendanceMonthYear: monthYear,
	})
	if err != nil {
		return "", err
	}
	return result, nil
}
