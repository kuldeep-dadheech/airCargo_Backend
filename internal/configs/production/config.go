package production

import (
	"aircargo/internal/configs"
	"aircargo/pkg/logging"
	"fmt"
	"os"
	"strconv"
)

func GetConfig(appName configs.AppName) configs.Config {

	var apiPort int = 80
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

	pg_port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return configs.Config{
		AppName:          appName,
		Port:             apiPort,
		GrpcPort:         5051,
		BaseRouterPrefix: "/data",
		RedisConfig: configs.RedisConfig{
			DSN:      os.Getenv("REDIS_DSN"),
			Password: &redisPassword,
			DB:       redisDb,
		},
		LogConfig: configs.LogConfig{
			LogSink:  configs.CONSOLE,
			LogLevel: logging.WARN,
		},
		PgDbConfig: configs.PgDbConfig{
			Host:           os.Getenv("PG_HOST"),
			Port:           pg_port,
			Username:       os.Getenv("PG_USERNAME"),
			Password:       os.Getenv("PG_PASSWORD"),
			Database:       os.Getenv("PG_DATABASE"),
			DatabasePrefix: os.Getenv("PG_DATABASE_PREFIX"),
			MaxConnections: 5,
			MaxIdle:        1,
		},
	}
}
