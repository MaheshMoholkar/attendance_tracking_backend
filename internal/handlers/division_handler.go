package handlers

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type DivisionHandler struct {
	store *database.Store
}

func NewDivisionHandler(store *database.Store) *DivisionHandler {
	return &DivisionHandler{
		store: store,
	}
}

func (h *DivisionHandler) HandleGetDivisions(ctx *fiber.Ctx) error {
	divisions, err := h.store.DB.GetDivisions(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(divisions)
}
