package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

var (
	SERVER_HTTP_PORT    string
	SERVER_HTTP_ADDRESS string

	APP_DEBUG_MODE bool

	DB_DRY_RUN       bool
	DB_NAME          string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_HOST          string
	DB_PORT          string
	DB_SSL_MODE      string
	DB_RUN_MIGRATION bool
	DB_RUN_SEEDER    bool

	KAFKA_BROKER_URLS    []string
	KAFKA_CONSUMER_GROUP string
)

func init() {
	InitConfig()
	var err error

	DB_DRY_RUN = false

	SERVER_HTTP_PORT = viper.GetString("app.port")
	DB_NAME = viper.GetString("db.name")
	DB_USERNAME = viper.GetString("db.username")
	DB_PASSWORD = viper.GetString("db.password")
	DB_HOST = viper.GetString("db.host")
	DB_PORT = viper.GetString("db.port")
	DB_SSL_MODE = viper.GetString("db.ssl_mode")

	DB_RUN_MIGRATION, err = strconv.ParseBool(viper.GetString("db.run_migrations"))
	if err != nil {
		DB_RUN_MIGRATION = false
	}

	DB_RUN_SEEDER, err = strconv.ParseBool(viper.GetString("db.run_seeders"))
	if err != nil {
		DB_RUN_SEEDER = false
	}

	KAFKA_CONSUMER_GROUP = viper.GetString("kafka.consumer_group")
	KAFKA_BROKER_URLS = viper.GetStringSlice("kafka.broker_urls")

}

func findConfigFile() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for {
		configPath := filepath.Join(cwd, "config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			break
		}
		cwd = parentDir
	}

	return "", fmt.Errorf("config file not found")
}

func InitConfig() error {
	// Find the configuration file
	configPath, err := findConfigFile()
	if err != nil {
		return err
	}
	// Set the config file
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// fmt.Println("Config loaded successfully")
	return nil
}

type FiberHttpServiceParams struct {
	Port    string
	Address string
}

func ProvideFiberHttpServiceParams() *FiberHttpServiceParams {

	fmt.Println("Port = ", SERVER_HTTP_PORT)

	return &FiberHttpServiceParams{
		Port:    SERVER_HTTP_PORT,
		Address: SERVER_HTTP_ADDRESS,
	}
}

func InitializeHTTPService(params *FiberHttpServiceParams) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:   false,
		BodyLimit: 200 * 1024 * 1024,
	})
	app.Use(fiberLog.New())
	// storageservice.NewS3Client(configs.S3REGION)
	// app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "*",
	}))

	log.Printf("Starting Fiber HTTP listener at Port [%s]...", params.Port)

	return app
}
