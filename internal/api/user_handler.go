package api

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *database.Store
}

func NewUserHandler(store *database.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandleCreateUser(ctx *fiber.Ctx) error {
	var params types.PostUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	return nil
}
