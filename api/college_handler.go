package api

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ClassQueryParams struct {
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

func (h *CollegeHandler) HandleGetClassInfo(ctx *fiber.Ctx) error {
	var qparams ClassQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}

	filter := bson.M{}

	classInfos, err := h.CollegeStore.GetClassInfo(ctx.Context(), filter)
	if err != nil {
		return err
	}

	return ctx.JSON(classInfos)
}

func (h *CollegeHandler) HandleGetClass(ctx *fiber.Ctx) error {
	var qparams ClassQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}
	filter := bson.M{}
	classes, err := h.CollegeStore.GetClasses(ctx.Context(), filter)
	if err != nil {
		return err
	}
	classNames := make([]string, len(classes))
	for i, class := range classes {
		classNames[i] = class.ClassName
	}
	return ctx.JSON(classNames)

}

func (h *CollegeHandler) HandlePostClass(ctx *fiber.Ctx) error {
	var params types.PostClassParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	class, err := types.NewClassFromParams(params)
	if err != nil {
		return err
	}
	insertedClass, err := h.CollegeStore.PostClass(ctx.Context(), class)
	if err != nil {
		return err
	}
	return ctx.JSON(insertedClass)
}
