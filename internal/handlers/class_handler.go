package handlers

import (
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	store *database.Store
}

func NewClassHandler(store *database.Store) *ClassHandler {
	return &ClassHandler{
		store: store,
	}
}

func (h *ClassHandler) HandleGetClasses(ctx *fiber.Ctx) error {
	classes, err := h.store.DB.GetClasses(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(classes)
}

func (h *ClassHandler) HandleCreateClass(ctx *fiber.Ctx) error {
	var class types.ClassInfo
	if err := ctx.BodyParser(&class); err != nil {
		return err
	}
	newClass, err := h.store.DB.CreateClassInfo(ctx.Context(), class.ClassName)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(newClass)
}

func (h *ClassHandler) HandleDeleteClass(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("class_id"))
	if err != nil {
		return err
	}
	err = h.store.DB.DeleteClassInfo(ctx.Context(), int32(id))
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
