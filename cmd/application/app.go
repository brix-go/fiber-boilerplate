package application

import (
	"fmt"
	"github.com/brix-go/fiber/config"
	redis_client "github.com/brix-go/fiber/infrastructure/Redis"
	"github.com/brix-go/fiber/infrastructure/database"
	userController "github.com/brix-go/fiber/internal/domain/user/controller"
	userRepository "github.com/brix-go/fiber/internal/domain/user/repository"
	userService "github.com/brix-go/fiber/internal/domain/user/service"
	middleware "github.com/brix-go/fiber/middleware/error"
	"github.com/brix-go/fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"log"
)

func Run() {
	config.LoadConfig()
	err := middleware.LoadErrorListFromJsonFile(config.AppConfig.ErrorContract.JSONPathFile)
	if err != nil {
		log.Fatal("Failed to read to errorContract.json:", err)
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	redisClient := redis_client.RedisClient

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(cors.New())

	app.Use(requestid.New(requestid.Config{
		Generator:  utils.UUIDv4,
		ContextKey: "request-id",
	}))

	//Todo : Define Repository here
	userRepo := userRepository.NewRepository(db)

	//Todo : Define Service here
	userSvc := userService.NewService(userRepo, redisClient)

	//Todo: Define controller
	userCtrl := userController.NewController(userSvc)

	routerApp := router.NewRouter(&router.RouteParams{
		userCtrl,
	})
	routerApp.SetupRoute(app)
	err = app.Listen(fmt.Sprintf(":%s", config.AppConfig.AppConfig.Port))
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
