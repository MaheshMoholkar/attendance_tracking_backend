package api

import (
	"log"
	"strconv"

	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/MaheshMoholkar/attendance_tracking_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type StudentQueryParams struct {
	rollno int
}

type StudentHandler struct {
	Store *db.Store
}

func NewStudentHandler(store *db.Store) *StudentHandler {
	return &StudentHandler{
		Store: store,
	}
}

func (h *StudentHandler) HandleGetStudents(ctx *fiber.Ctx) error {
	var qparams StudentQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}
	var filter bson.M
	if qparams.rollno != 0 {
		filter = bson.M{"rollno": qparams.rollno}
	}
	students, err := h.Store.StudentStore.GetStudents(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return ctx.JSON(students)
}

func (h *StudentHandler) HandleCreateStudent(ctx *fiber.Ctx) error {
	var params types.PostStudentParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	student, err := types.NewStudentFromParams(params)
	if err != nil {
		return err
	}
	insertedStudent, err := h.Store.StudentStore.PostStudent(ctx.Context(), student)
	if err != nil {
		return err
	}
	// Retrieve the class based on className
	filter := bson.M{"className": params.ClassName}
	class, err := h.Store.CollegeStore.GetClassByName(ctx.Context(), filter)
	if err != nil {
		return err
	}

	// Add the student ID to the division
	class.Divisions[params.Division] = append(class.Divisions[params.Division], insertedStudent.ID)

	// Update the class document in the database
	update := bson.M{"$set": bson.M{"divisions": class.Divisions}}
	_, err = h.Store.CollegeStore.UpdateClass(ctx.Context(), filter, update)
	if err != nil {
		return err
	}

	return ctx.JSON(insertedStudent)
}

func (h *StudentHandler) HandleDeleteStudent(ctx *fiber.Ctx) error {
	rollnoStr := ctx.Query("rollno")
	if rollnoStr == "" {
		log.Println("Rollno parameter is required")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Rollno parameter is required",
		})
	}

	rollno, err := strconv.Atoi(rollnoStr)
	if err != nil {
		log.Printf("Invalid rollno parameter: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid rollno parameter",
		})
	}
	filter := bson.M{"rollno": rollno}
	err = h.Store.StudentStore.DeleteStudent(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return nil

}
