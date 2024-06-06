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

		// Fetch className and divisionName using student_id from student_info table
		studentInfo, err := h.store.DB.GetStudentInfo(c.Context(), studentID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get student info")
		}
		className := studentInfo.Classname
		divisionName := studentInfo.Division

		// Get class_id from class_info table
		classID, err := h.store.DB.GetClassIDByName(c.Context(), className)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get class ID")
		}

		// Get division_id from division_info table using class_id
		divisionID, err := h.store.DB.GetDivisionIDByNameAndClass(c.Context(), postgres.GetDivisionIDByNameAndClassParams{
			Divisionname: divisionName,
			ClassID:      classID,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get division ID")
		}

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
	className := ctx.Query("class_name")
	divisionName := ctx.Query("division_name")
	dateStr := ctx.Query("date")

	// Parse the date
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid date format")
	}

	// Get class ID
	classID, err := h.store.DB.GetClassIDByName(ctx.Context(), className)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Class not found")
	}

	// Get division ID
	divisionID, err := h.store.DB.GetDivisionIDByNameAndClass(ctx.Context(), postgres.GetDivisionIDByNameAndClassParams{
		Divisionname: divisionName,
		ClassID:      classID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Division not found")
	}

	// Get attendance
	attendances, err := h.store.DB.GetAttendanceByClassDivisionAndMonthYear(ctx.Context(), postgres.GetAttendanceByClassDivisionAndMonthYearParams{
		ClassID:    classID,
		DivisionID: divisionID,
		Date:       date,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error fetching attendance data")
	}

	// Convert postgres.AttendanceInfo to types.Attendance
	convertedAttendances := types.ParseAttendances(attendances)

	// Convert attendance data to the desired format
	result := types.ConvertAttendanceData(convertedAttendances)

	return ctx.JSON(result)
}
