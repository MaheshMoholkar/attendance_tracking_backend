package api

import (
	"errors"

	"github.com/MaheshMoholkar/attendance_tracking_backend/api/middleware"
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

func (h *AuthHandler) HandleLogin(ctx *fiber.Ctx) error {
	var params AuthParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	id, err := h.userStore.GetUserByEmail(ctx.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.JSON(map[string]string{"msg": "Not Found"})
		}
		return err
	}
	token, err := middleware.GenerateJWT(id)
	if err != nil {
		return err
	}
	return ctx.JSON(token)
}
