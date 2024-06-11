package handlers

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct {
	store *database.Store
}

func NewSubjectHandler(store *database.Store) *SubjectHandler {
	return &SubjectHandler{
		store: store,
	}
}

func (h *SubjectHandler) HangeGetSubjects(ctx *fiber.Ctx) error {
	subjects, err := h.store.DB.GetSubjects(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(subjects)
}

func (h *SubjectHandler) HandleCreateSubject(ctx *fiber.Ctx) error {
	var subject types.Subject
	if err := ctx.BodyParser(&subject); err != nil {
		return err
	}
	_, err := h.store.DB.CreateSubjectInfo(ctx.Context(), postgres.CreateSubjectInfoParams{
		Subjectname: subject.SubjectName,
		ClassID:     subject.ClassID,
	})
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusCreated)
}
