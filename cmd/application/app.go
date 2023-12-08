package application

import (
	"fmt"
	"github.com/brix-go/fiber/config"
	redis_client "github.com/brix-go/fiber/infrastructure/Redis"
	"github.com/brix-go/fiber/infrastructure/database"
	"github.com/brix-go/fiber/infrastructure/kafka"
	infrastructure "github.com/brix-go/fiber/infrastructure/log"
	userController "github.com/brix-go/fiber/internal/domain/user/controller"
	userRepository "github.com/brix-go/fiber/internal/domain/user/repository"
	userService "github.com/brix-go/fiber/internal/domain/user/service"
	middleware "github.com/brix-go/fiber/middleware/error"
	middlewareLog "github.com/brix-go/fiber/middleware/log"
	"github.com/brix-go/fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func Run() {
	config.LoadConfig()
	log := infrastructure.NewLogCustom()
	err := middleware.LoadErrorListFromJsonFile(config.AppConfig.ErrorContract.JSONPathFile)
	if err != nil {
		log.Logrus.Fatal("Failed to read to errorContract.json:", err)
	}

	// Database
	db := database.ConnectDatabase()

	// Redis
	redisClient := redis_client.RedisClient

	// Kafka
	_ = kafka.NewKafkaConsumer(*log)
	_ = kafka.NewKafkaProducer(*log)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Jakarta",
		Done: func(ctx *fiber.Ctx, logString []byte) {
			middlewareLog.LogMiddleware(ctx, logString, log)
		},
	}))

	app.Use(requestid.New(requestid.Config{
		Generator:  utils.UUIDv4,
		ContextKey: "request-id",
	}))

	//Todo : Define Repository here
	redisRepo := redis_client.NewRedisRepository(redisClient)
	userRepo := userRepository.NewRepository()

	//Todo : Define Service here
	userSvc := userService.NewService(db.DB, userRepo, &redisRepo)

	//Todo: Define controller
	userCtrl := userController.NewController(userSvc, log)

	routerApp := router.NewRouter(&router.RouteParams{
		userCtrl,
	})
	routerApp.SetupRoute(app)
	err = app.Listen(fmt.Sprintf(":%s", config.AppConfig.AppConfig.Port))
	if err != nil {
		log.Logrus.Fatal("Failed to start server:", err)
	}
}
