package main

import (
	"embed"
	"flag"
	"os"
	"sagebackend/cmd/api"
	"sagebackend/cmd/migrations"
	"sagebackend/cmd/rpc"
	"sagebackend/internal/configs"
	"sagebackend/internal/configs/local"
	"sagebackend/internal/configs/production"
	"sagebackend/internal/configs/sit"
	"sagebackend/internal/configs/uat"

	"github.com/go-redis/redis/v7"
	// "github.com/gomodule/redigo/redis"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS
var Client *redis.Client

const (
	API        configs.AppName = "sage-backend-http-api"
	RPC        configs.AppName = "sage-backend-rpc-server"
	MIGRATIONS configs.AppName = "sage-backend-migrations"
)

func main() {

	program := flag.String("program", "", "The Program that needs to run")
	envTier := os.Getenv("TIER")

	flag.Parse()
	// initialsize redis :
	switch *program {
	case "migrations":
		config := getConfig(envTier, MIGRATIONS)
		migrations.Initialize(embedMigrations, config)
	case "http-api":
		config := getConfig(envTier, API)
		api.Initialize(config)
	case "rpc-server":
		config := getConfig(envTier, RPC)
		rpc.Initialize(config)
	default:
		panic("could not understand the program that needed to be run")
	}
}

func getConfig(
	envTier string,
	appName configs.AppName,
) configs.Config {
	var config configs.Config

	switch envTier {
	case string(configs.LOCAL):
		config = local.GetConfig(appName)
	case string(configs.UAT):
		config = uat.GetConfig(appName)
	case string(configs.SIT):
		config = sit.GetConfig(appName)
	case string(configs.PRODUCTION):
		config = production.GetConfig(appName)
	default:
		panic("invalid tier was not understood by the system")
	}

	return config
}
