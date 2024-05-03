package configs

import (
	"strconv"

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
