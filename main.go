package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/handlers"
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {
	godotenv.Load(".env")

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	listenAddr := flag.String("listenAddr", address, "The listen address of the api server")
	flag.Parse()

	dbQueries, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	store := database.New(dbQueries)

	var (
		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")

		// Initialize handlers
		authHandler       = handlers.NewAuthHandler(store)
		studentHandler    = handlers.NewStudentHandler(store)
		staffHandler      = handlers.NewStaffHandler(store)
		classHandler      = handlers.NewClassHandler(store)
		divisionHandler   = handlers.NewDivisionHandler(store)
		attendanceHandler = handlers.NewAttendanceHandler(store)
	)

	// Customize the CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Middlewares
	apiv1.Use(middleware.Logger)

	// Auth handlers
	app.Post("/api/auth/register", staffHandler.HandleCreateStaff)
	app.Post("/api/auth/login", authHandler.HandleUserLogin)

	// Staff handlers
	apiv1.Get("/staff", staffHandler.HandleGetStaff)
	apiv1.Get("/staffs", staffHandler.HandleGetStaffs)
	apiv1.Put("/staff", staffHandler.HandleUpdateStaff)
	apiv1.Delete("/staff", staffHandler.HandleDeleteStaff)

	// Student handlers
	apiv1.Get("/student", studentHandler.HandleGetStudent)
	apiv1.Get("/students", studentHandler.HandleGetStudents)
	apiv1.Post("/student", studentHandler.HandleCreateStudent)
	apiv1.Put("/student", studentHandler.HandleUpdateStudent)
	apiv1.Delete("/student", studentHandler.HandleDeleteStudent)

	// Class handlers
	apiv1.Get("/classes", classHandler.HandleGetClasses)
	apiv1.Post("/class", classHandler.HandleCreateClass)
	apiv1.Delete("/class", classHandler.HandleDeleteClass)
	apiv1.Get("/class-divisions", classHandler.HandleGetClassDivisions)

	// Division handlers
	apiv1.Get("/divisions", divisionHandler.HandleGetDivisions)
	apiv1.Post("/division", divisionHandler.HandleCreateDivision)
	apiv1.Delete("/division", divisionHandler.HandleDeleteDivision)

	// Attendance handlers

	apiv1.Get("/attendance", attendanceHandler.InitializeAttendanceTableHandler)
	apiv1.Post("/attendance", attendanceHandler.UpdateAttendanceHandler)

	app.Listen(*listenAddr)
}
