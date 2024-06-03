package handlers

import (
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type StudentHandler struct {
	store *database.Store
}

func NewStudentHandler(store *database.Store) *StudentHandler {
	return &StudentHandler{
		store: store,
	}
}

func (h *StudentHandler) HandleGetStudents(ctx *fiber.Ctx) error {
	students, err := h.store.DB.GetStudentsInfo(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(students)
}

func (h *StudentHandler) HandleCreateStudent(ctx *fiber.Ctx) error {
	var params types.Student
	err := ctx.BodyParser(&params)
	if err != nil {
		return err
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(int(params.StudentID))), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	passwordHash := string(hashedBytes)

	_, err = h.store.DB.CreateStudentCredentials(ctx.Context(), postgres.CreateStudentCredentialsParams{
		StudentID:    int32(params.StudentID),
		PasswordHash: passwordHash,
	})
	if err != nil {
		return err
	}

	_, err = h.store.DB.CreateStudentInfo(ctx.Context(), postgres.CreateStudentInfoParams{
		Firstname: params.FirstName,
		Lastname:  params.LastName,
		Rollno:    int32(params.Rollno),
		Email:     params.Email,
		Classname: params.ClassName,
		Division:  params.Division,
		Year:      int32(params.Year),
		StudentID: int32(params.StudentID),
	})
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *StudentHandler) HandleUpdateStudent(ctx *fiber.Ctx) error {
	var params types.Student
	err := ctx.BodyParser(&params)
	if err != nil {
		return err
	}

	_, err = h.store.DB.UpdateStudentInfo(ctx.Context(), postgres.UpdateStudentInfoParams{
		Firstname: params.FirstName,
		Lastname:  params.LastName,
		Rollno:    int32(params.Rollno),
		Email:     params.Email,
		Classname: params.ClassName,
		Division:  params.Division,
		Year:      int32(params.Year),
		StudentID: int32(params.StudentID),
	})
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *StudentHandler) HandleDeleteStudent(ctx *fiber.Ctx) error {
	studentIDStr := ctx.Query("student_id")
	if studentIDStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "student_id parameter is required")
	}

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid student_id parameter")
	}

	_, err = h.store.DB.DeleteStudentInfo(ctx.Context(), int32(studentID))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid student_id")
	}
	err = h.store.DB.DeleteStudentCredentials(ctx.Context(), int32(studentID))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid student_id")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
