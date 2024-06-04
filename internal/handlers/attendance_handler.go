package handlers

import (
	"strconv"
	"time"

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

func (h *AttendanceHandler) HandleInsertAttendance(c *fiber.Ctx) error {
	var req types.AttendanceList
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	for _, attendanceMap := range req.Data {
		studentID := attendanceMap.StudentID
		classID := attendanceMap.ClassID
		divisionID := attendanceMap.DivisionID

		for day, status := range attendanceMap.Attendance {
			dayInt, err := strconv.Atoi(day)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid day in attendance data")
			}
			date := time.Date(req.Year, time.Month(req.Month), dayInt, 0, 0, 0, 0, time.UTC)

			err = h.store.DB.InsertAttendance(c.Context(), postgres.InsertAttendanceParams{
				StudentID:  studentID,
				Date:       date,
				Status:     status.(bool),
				ClassID:    classID,
				DivisionID: divisionID,
			})
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Failed to save attendance")
			}
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *AttendanceHandler) HandleGetAttendanceByStudent(ctx *fiber.Ctx) error {
	studentID, err := strconv.Atoi(ctx.Params("student_id"))
	if err != nil {
		return err
	}

	attendances, err := h.store.DB.GetAttendanceByStudent(ctx.Context(), int32(studentID))
	if err != nil {
		return err
	}
	return ctx.JSON(attendances)
}

func (h *AttendanceHandler) HandleGetAttendanceList(ctx *fiber.Ctx) error {
	classID, err := strconv.Atoi(ctx.Query("class_id"))
	if err != nil {
		return err
	}
	divisionID, err := strconv.Atoi(ctx.Query("division_id"))
	if err != nil {
		return err
	}
	dateStr := ctx.Query("date") // Format: YYYY-MM
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		return err
	}

	attendances, err := h.store.DB.GetAttendanceByClassDivisionAndMonthYear(ctx.Context(), postgres.GetAttendanceByClassDivisionAndMonthYearParams{
		ClassID:    int32(classID),
		DivisionID: int32(divisionID),
		Date:       date,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(attendances)
}
