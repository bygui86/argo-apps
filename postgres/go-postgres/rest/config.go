package rest

import (
	"github.com/bygui86/go-postgres/logging"
	"github.com/bygui86/go-postgres/utils"
)

const (
	dbHostEnvVar     = "DB_HOST"
	dbPortEnvVar     = "DB_PORT"
	dbUsernameEnvVar = "DB_USERNAME"
	dbPasswordEnvVar = "DB_PASSWORD"
	dbNameEnvVar     = "DB_NAME"
	dbSslModeEnvVar  = "DB_SSL_MODE"
	restHostEnvVar   = "REST_HOST"
	restPortEnvVar   = "REST_PORT"

	dbHostEnvVarDefault     = "localhost"
	dbPortEnvVarDefault     = 5432
	dbUsernameEnvVarDefault = "username"
	dbPasswordEnvVarDefault = "password"
	dbNameEnvVarDefault     = "db"
	dbSslModeEnvVarDefault  = "disable"
	restHostEnvVarDefault   = "localhost"
	restPortEnvVarDefault   = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load configurations")
	return &config{
		DbHost:     utils.GetStringEnv(dbHostEnvVar, dbHostEnvVarDefault),
		DbPort:     utils.GetIntEnv(dbPortEnvVar, dbPortEnvVarDefault),
		DbUsername: utils.GetStringEnv(dbUsernameEnvVar, dbUsernameEnvVarDefault),
		DbPassword: utils.GetStringEnv(dbPasswordEnvVar, dbPasswordEnvVarDefault),
		DbName:     utils.GetStringEnv(dbNameEnvVar, dbNameEnvVarDefault),
		DbSslMode:  utils.GetStringEnv(dbSslModeEnvVar, dbSslModeEnvVarDefault),
		RestHost:   utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		RestPort:   utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
