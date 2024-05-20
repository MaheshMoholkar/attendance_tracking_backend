package api

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type StudentQueryParams struct {
	rollno int
}

type StudentHandler struct {
	StudentStore db.StudentStore
}

func NewStudentHandler(studentStore db.StudentStore) *StudentHandler {
	return &StudentHandler{
		StudentStore: studentStore,
	}
}

func (h *StudentHandler) HandleGetStudents(ctx *fiber.Ctx) error {
	var qparams StudentQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}
	var filter bson.M
	if qparams.rollno != 0 {
		filter = bson.M{"rollno": qparams.rollno}
	}
	students, err := h.StudentStore.GetStudents(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return ctx.JSON(students)
}

func (h *StudentHandler) HandleCreateStudent(ctx *fiber.Ctx) error {
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
