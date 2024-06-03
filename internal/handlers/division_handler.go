package handlers

import (
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
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

func (h *DivisionHandler) HandleCreateDivision(ctx *fiber.Ctx) error {
	var division types.Division
	if err := ctx.BodyParser(&division); err != nil {
		return err
	}
	_, err := h.store.DB.CreateDivisionInfo(ctx.Context(), postgres.CreateDivisionInfoParams{
		Divisionname: division.DivisionName,
		ClassID:      division.ClassID,
	})
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

func (h *DivisionHandler) HandleDeleteDivision(ctx *fiber.Ctx) error {
	id := ctx.Query("division_id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "division_id is required")
	}

	division_id, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid division_id")
	}
	err = h.store.DB.DeleteDivisionInfo(ctx.Context(), int32(division_id))
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
