package local

import (
	"aircargo/internal/configs"
	"aircargo/pkg/logging"
	"fmt"
	"os"
	"strconv"
)

func GetConfig(appName configs.AppName) configs.Config {

	pgPort, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	var apiPort int = 9000
	if len(os.Getenv("HTTP_API_PORT")) > 0 {
		p, err := strconv.Atoi(os.Getenv("HTTP_API_PORT"))

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		} else {
			apiPort = p
		}
	}

	var grpcPort int = 5051
	if len(os.Getenv("GRPC_PORT")) > 0 {
		p, err := strconv.Atoi(os.Getenv("GRPC_PORT"))

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		} else {
			grpcPort = p
		}
	}
	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")

	return configs.Config{
		AppName:          appName,
		Port:             apiPort,
		GrpcPort:         grpcPort,
		BaseRouterPrefix: "",
		RedisConfig: configs.RedisConfig{
			DSN:      os.Getenv("REDIS_DSN"),
			Password: &redisPassword,
			DB:       redisDb,
		},
		EnableAPM: false,
		LogConfig: configs.LogConfig{
			LogSink:  configs.CONSOLE,
			LogLevel: logging.INFO,
		},
		PgDbConfig: configs.PgDbConfig{
			Host:           os.Getenv("PG_HOST"),
			Port:           pgPort,
			Username:       os.Getenv("PG_USERNAME"),
			Password:       os.Getenv("PG_PASSWORD"),
			Database:       os.Getenv("PG_DATABASE"),
			DatabasePrefix: os.Getenv("PG_DATABASE_PREFIX"),
			MaxConnections: 5,
			MaxIdle:        1,
			SSLMode:        configs.SSL_MODE_DISABLED,
		},
		EmailConfig: configs.EmailConfig{
			Email:             "no-reply@wizfreight.com",
			ServerToken:       os.Getenv("POSTMARK_SERVER_TOKEN"),
			MesaircargoStream: os.Getenv("POSTMARK_MESaircargo_STREAM"),
		},
	}
}
