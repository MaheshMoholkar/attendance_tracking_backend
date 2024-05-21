package api

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type AttendanceQueryParams struct {
	rollno int
}

type AttendanceHandler struct {
	AttendanceStore db.AttendanceStore
}

func NewAttendanceHandler(attendanceStore db.AttendanceStore) *AttendanceHandler {
	return &AttendanceHandler{
		AttendanceStore: attendanceStore,
	}
}

func (h *AttendanceHandler) HandleGetAttendance(ctx *fiber.Ctx) error {
	var qparams AttendanceQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}
	var filter bson.M
	if qparams.rollno != 0 {
		filter = bson.M{"rollno": qparams.rollno}
	}
	attendance, err := h.AttendanceStore.GetAttendance(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return ctx.JSON(attendance)
}

func (h *AttendanceHandler) HandlePostAttendance(ctx *fiber.Ctx) error {
	var params types.PostAttendanceParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	attendance, err := types.NewAttendaceFromParams(params)
	if err != nil {
		return err
	}
	insertedAttedance, err := h.AttendanceStore.PostAttendance(ctx.Context(), attendance)
	if err != nil {
		return err
	}
	return ctx.JSON(insertedAttedance)
}
