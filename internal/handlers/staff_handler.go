package handlers

import (
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type StaffHandler struct {
	store *database.Store
}

func NewStaffHandler(store *database.Store) *StaffHandler {
	return &StaffHandler{
		store: store,
	}
}

func (h *StaffHandler) HandleGetStaff(ctx *fiber.Ctx) error {
	staffIDStr := ctx.Query("staff_id")
	if staffIDStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "staff_id parameter is required")
	}

	staffID, err := strconv.Atoi(staffIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid staff_id parameter")
	}

	staff, err := h.store.DB.GetStaffInfo(ctx.Context(), int32(staffID))
	if err != nil {
		return err
	}
	return ctx.JSON(staff)
}

func (h *StaffHandler) HandleGetStaffs(ctx *fiber.Ctx) error {
	staffs, err := h.store.DB.GetStaffsInfo(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(staffs)
}

func (h *StaffHandler) HandleCreateStaff(ctx *fiber.Ctx) error {
	var params types.Staff
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create staff credentials
	if _, err := h.store.DB.CreateStaffCredentials(ctx.Context(), postgres.CreateStaffCredentialsParams{
		StaffID:      params.StaffID,
		PasswordHash: string(hashedPassword),
	}); err != nil {
		return err
	}

	// Create staff info
	if _, err := h.store.DB.CreateStaffInfo(ctx.Context(), postgres.CreateStaffInfoParams{
		Firstname: params.FirstName,
		Lastname:  params.LastName,
		Email:     params.Email,
		StaffID:   params.StaffID,
	}); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *StaffHandler) HandleUpdateStaff(ctx *fiber.Ctx) error {
	var params types.Staff
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	// Update staff info
	if _, err := h.store.DB.UpdateStaffInfo(ctx.Context(), postgres.UpdateStaffInfoParams{
		Firstname: params.FirstName,
		Lastname:  params.LastName,
		Email:     params.Email,
		StaffID:   params.StaffID,
	}); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *StaffHandler) HandleDeleteStaff(ctx *fiber.Ctx) error {
	staffIDStr := ctx.Query("staff_id")
	if staffIDStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "staff_id parameter is required")
	}

	staffID, err := strconv.Atoi(staffIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid staff_id parameter")
	}

	// Delete staff credentials
	if err := h.store.DB.DeleteStaffCredentials(ctx.Context(), int32(staffID)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid staff_id")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
