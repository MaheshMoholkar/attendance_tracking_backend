package main

import (
	"log"
	"os"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/api"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {
	listenAddr := os.Getenv("PORT")

	dbQueries, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	store := database.New(dbQueries)

	var (
		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")

		// Initialize handlers
		authHandler       = api.NewAuthHandler(store)
		studentHandler    = api.NewStudentHandler(store)
		userHandler       = api.NewUserHandler(store)
		attendanceHandler = api.NewAttendanceHandler(store)
	)

	// Allow CORS
	app.Use(cors.New())

	// Or customize the configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Middlewares
	apiv1.Use(middleware.Logger)

	// Auth handlers
	app.Post("/api/auth/register", userHandler.HandleCreateUser)
	app.Post("/api/auth/login", authHandler.HandleUserLogin)

	// Student handlers
	apiv1.Get("/students", studentHandler.HandleGetStudents)
	apiv1.Post("/student", studentHandler.HandleCreateStudent)
	apiv1.Put("/student", studentHandler.HandleUpdateStudent)

	// Attendance handlers
	apiv1.Get("/attendance", attendanceHandler.HandleGetAttendance)
	apiv1.Post("/attendance", attendanceHandler.HandlePostAttendance)

	app.Listen(listenAddr)
}
