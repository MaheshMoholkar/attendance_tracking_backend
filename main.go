package main

import (
	"context"
	"flag"
	"log"

	"github.com/MaheshMoholkar/attendance_tracking_backend/api"
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DB_URI))
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	var (
		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")

		//initialize stores
		userStore       = db.NewMongoUserStore(client)
		studentStore    = db.NewMongoStudentStore(client)
		attendanceStore = db.NewMongoAttendanceStore(client)
		collegeStore    = db.NewMongoCollegeStore(client)
		_               = &db.Store{
			UserStore: userStore,
		}

		// initialize handlers
		authHandler       = api.NewAuthHandler(userStore)
		userHandler       = api.NewUserHandler(userStore)
		studentHandler    = api.NewStudentHandler(studentStore)
		attendanceHandler = api.NewAttendanceHandler(attendanceStore)
		collegeHandler    = api.NewCollegeHandler(collegeStore)
	)

	// allow cors
	app.Use(cors.New())

	// Or customize the configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// middlewares
	//apiv1.Use(middleware.VerifyToken())

	// auth handlers
	app.Post("/api/auth/register", userHandler.HandleCreateUser)
	app.Post("/api/auth/login", authHandler.HandleLogin)

	// user handlers
	apiv1.Get("/classes", collegeHandler.HandleGetClasses)
	apiv1.Post("/class", collegeHandler.HandlePostClass)

	// student handlers
	apiv1.Get("/students", studentHandler.HandleGetStudents)
	apiv1.Post("/student", studentHandler.HandleCreateStudent)

	// attendance handlers
	apiv1.Get("/attendance", attendanceHandler.HandleGetAttendance)
	apiv1.Post("/attendance", attendanceHandler.HandlePostAttendance)

	app.Listen(*listenAddr)
}
