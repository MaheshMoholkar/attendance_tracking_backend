package api

import (
	"log"

	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	var params types.PostUserParams
	if err := ctx.BodyParser(&params); err != nil {
		return nil
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.PostUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	log.Print(insertedUser)
	return ctx.JSON(insertedUser)
}
