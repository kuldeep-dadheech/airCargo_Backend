package configs

import (
	"sagebackend/pkg/logging"
)

type AppPort uint64
type AppName string
type BaseRouterPrefix string

type Config struct {
	AppName          AppName
	Port             int
	GrpcPort         int
	EnableAPM        bool
	BaseRouterPrefix BaseRouterPrefix
	PgDbConfig       PgDbConfig
	RedisConfig      RedisConfig
	LogConfig        LogConfig
	EmailConfig 	 EmailConfig
}

type PgSSLMode string

const (
	SSL_MODE_REQUIRED PgSSLMode = "require"
	SSL_MODE_DISABLED PgSSLMode = "disable"
)

type PgDbConfig struct {
	Host           string
	Port           int
	Username       string
	Password       string
	Database       string
	DatabasePrefix string
	MaxConnections int
	MaxIdle        int
	SSLMode        PgSSLMode
}

type RedisConfig struct {
	DSN      string
	Password *string
	DB       int
}
type LogSink int

const (
	CONSOLE LogSink = iota
	STDOUT  LogSink = iota
)

type LogConfig struct {
	LogSink  LogSink
	LogLevel logging.LogLevel
}

type ConfigTier string

const (
	PRODUCTION ConfigTier = "PRODUCTION"
	SIT        ConfigTier = "SIT"
	LOCAL      ConfigTier = "LOCAL"
	UAT        ConfigTier = "UAT"
)

type EmailConfig struct {
	Email string
	ServerToken string
	MessageStream string
}