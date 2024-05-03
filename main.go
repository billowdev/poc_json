package main

import (
	"fmt"
	"poc_json/configs"
	"poc_json/handlers"
	"poc_json/models"
	"poc_json/repositories"
	"poc_json/services"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		configs.DB_HOST,
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_NAME,
		configs.DB_PORT,
	)
	ConDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	if configs.DB_RUN_MIGRATION {
		if err := models.RunMigrations(ConDB); err != nil {
			panic("failed to run migration")
		}
	}
	if configs.DB_RUN_SEEDER {
		models.RunSeeds(ConDB)
	}
	params := configs.ProvideFiberHttpServiceParams()
	fiberService := configs.InitializeHTTPService(params)
	producer := configs.InitSaramaProducer()
	kafkaConsumerGroup := configs.InitSaramaConsumer()
	notificationRepo := repositories.NewNotiRepo(ConDB)
	notificationService := services.NewNotiSrv(notificationRepo)
	notiHandler := handlers.NewNotiHandler(notificationService, kafkaConsumerGroup)

	topics := models.Topics
	notiHandler.StartConsuming(topics)

	fiberService = AppContainer(fiberService, ConDB, producer)
	portString := fmt.Sprintf(":%v", params.Port)

	err = fiberService.Listen(portString)

	if err != nil {
		panic("Failed to start golang Fiber server")
	}
}

type RouterDeps struct {
	route fiber.Router
}

func NewRoute(r fiber.Router) RouterDeps {
	return RouterDeps{r}
}

func AppContainer(app *fiber.App, db *gorm.DB, kafkaProducer *sarama.SyncProducer) *fiber.App {
	configurationEndpoint := app.Group("/api/v1")
	configurationRoute := NewRoute(configurationEndpoint)
	API(configurationRoute, db, kafkaProducer)

	return app
}

func API(route RouterDeps, db *gorm.DB, kafkaProducer *sarama.SyncProducer) {
	documentRepo := repositories.NewDocumentRepo(db)
	documentSrv := services.NewDocumentSrv(documentRepo, kafkaProducer)
	documentHandler := handlers.NewDocumentHandler(documentSrv)

	route.CreateDocumentRoute(documentHandler)
}

func (r RouterDeps) CreateDocumentRoute(h handlers.IDocumentHandlerInfs) {
	r.route.Post("/document", h.HandleTest)
}
