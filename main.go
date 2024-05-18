package main

import (
	"context"
	"flag"
	"log"

	"github.com/MaheshMoholkar/attendance_tracking_backend/api"
	"github.com/MaheshMoholkar/attendance_tracking_backend/api/middleware"
	"github.com/MaheshMoholkar/attendance_tracking_backend/db"
	"github.com/gofiber/fiber/v2"
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

	var (
		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")

		//initialize stores
		userStore    = db.NewMongoUserStore(client)
		studentStore = db.NewMongoStudentStore(client)
		_            = &db.Store{
			UserStore: userStore,
		}

		// initialize handlers
		authHandler    = api.NewAuthHandler(userStore)
		userHandler    = api.NewUserHandler(userStore)
		studentHandler = api.NewStudentHandler(studentStore)
	)

	// middlewares
	apiv1.Use(middleware.VerifyToken())

	// auth handlers
	app.Post("/api/auth/register", userHandler.HandleCreateUser)
	app.Post("/api/auth/login", authHandler.HandleLogin)

	apiv1.Post("student", studentHandler.HandleCreateStudent)

	app.Listen(*listenAddr)
}
