package handlers

import (
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthParams struct {
	StaffID  int    `json:"staff_id"`
	Password string `json:"password"`
}

type AuthHandler struct {
	store *database.Store
}

func NewAuthHandler(store *database.Store) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) HandleUserLogin(ctx *fiber.Ctx) error {
	var params AuthParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	passwordHash, err := h.store.DB.GetStaffCredentials(ctx.Context(), int32(params.StaffID))
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(params.Password))
	if err != nil {
		return err
	}

	token, err := middleware.GenerateJWT(strconv.Itoa(params.StaffID))
	if err != nil {
		return err
	}
	return ctx.JSON(token)
}
