package handlers

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
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
