package handlers

import (
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthParams struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"`
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
	err := ctx.BodyParser(&params)
	if err != nil {
		return err
	}

	var passwordHash string

	if params.Role == "staff" {
		passwordHash, err = h.store.DB.GetStaffCredentials(ctx.Context(), int32(params.ID))
		if err != nil {
			return err
		}
	} else {
		passwordHash, err = h.store.DB.GetStudentCredentials(ctx.Context(), int32(params.ID))
		if err != nil {
			return err
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(params.Password))
	if err != nil {
		return err

	}

	token, err := middleware.GenerateJWT(strconv.Itoa(params.ID), params.Role)
	if err != nil {
		return err
	}
	return ctx.JSON(token)
}
