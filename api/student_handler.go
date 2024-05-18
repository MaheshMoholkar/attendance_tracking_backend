package api

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	StudentStore db.StudentStore
}

func NewStudentHandler(studentStore db.StudentStore) *StudentHandler {
	return &StudentHandler{
		StudentStore: studentStore,
	}
}

func (h *StudentHandler) HandlerCreateStudent(ctx *fiber.Ctx) error {
	var params types.PostStudentParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	student, err := types.NewStudentFromParams(params)
	if err != nil {
		return err
	}
	insertedStudent, err := h.StudentStore.PostStudent(ctx.Context(), student)
	if err != nil {
		return err
	}
	return ctx.JSON(insertedStudent)
}
