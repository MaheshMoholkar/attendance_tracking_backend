package handlers

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type StaffHandler struct {
	store *database.Store
}

func NewUserHandler(store *database.Store) *StaffHandler {
	return &StaffHandler{
		store: store,
	}
}

func (h *StaffHandler) HandleCreateStaff(ctx *fiber.Ctx) error {
	var params types.PostUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}

	return nil
}
