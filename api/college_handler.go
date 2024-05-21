package api

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ClassesQueryParams struct {
	className string
}

type CollegeHandler struct {
	CollegeStore db.CollegeStore
}

func NewCollegeHandler(collegeStore db.CollegeStore) *CollegeHandler {
	return &CollegeHandler{
		CollegeStore: collegeStore,
	}
}

func (h *CollegeHandler) HandleGetClasses(ctx *fiber.Ctx) error {
	var qparams ClassesQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}
	filter := bson.M{"className": qparams.className}
	classes, err := h.CollegeStore.GetClasses(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return ctx.JSON(classes)

}
