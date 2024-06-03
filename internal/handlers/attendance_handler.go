package handlers

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	store *database.Store
}

func NewAttendanceHandler(store *database.Store) *AttendanceHandler {
	return &AttendanceHandler{
		store: store,
	}
}

func (s *AttendanceHandler) HandleGetAttendance(ctx *fiber.Ctx) error {
	return nil
}

func (h *AttendanceHandler) HandlePostAttendance(ctx *fiber.Ctx) error {
	return nil
}
