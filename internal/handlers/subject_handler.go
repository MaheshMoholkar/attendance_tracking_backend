package handlers

import (
	"strconv"

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

func (h *SubjectHandler) HandleDeleteSubject(ctx *fiber.Ctx) error {
	id := ctx.Query("subject_id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "subject_id is required")
	}

	subject_id, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid subject_id")
	}
	err = h.store.DB.DeleteSubjectInfo(ctx.Context(), int32(subject_id))
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
