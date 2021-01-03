package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ExpansiveWorlds/instrumentedsql"
	instrumentedsqlopentracing "github.com/ExpansiveWorlds/instrumentedsql/opentracing"
	"github.com/lib/pq"

	"github.com/bygui86/go-traces/http-server/logging"
)

const (
	// no tracing
	dbConnectionStringFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dbDriverName             = "postgres"

	// with tracing
	dbConnectionString       = "postgres://%s:%s@%s:%d/%s?sslmode=%s"
	instrumentedDbDriverName = "instrumeted-" + dbDriverName
)

func New() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector")

	cfg := loadConfig()

	db, dbErr := sql.Open(
		dbDriverName,
		fmt.Sprintf(dbConnectionStringFormat,
			cfg.dbHost, cfg.dbPort,
			cfg.dbUsername, cfg.dbPassword, cfg.dbName,
			cfg.dbSslMode,
		),
	)
	if dbErr != nil {
		return nil, dbErr
	}

	_, tableErr := db.Exec(createTableQuery)
	if tableErr != nil {
		return nil, tableErr
	}

	return db, nil
}

func NewWithWrappedTracing() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector with tracing")

	cfg := loadConfig()

	// Get a database driver.Connector for a fixed configuration.
	connector, connErr := pq.NewConnector(fmt.Sprintf(dbConnectionString,
		cfg.dbUsername, cfg.dbPassword,
		cfg.dbHost, cfg.dbPort,
		cfg.dbName, cfg.dbSslMode,
	))
	if connErr != nil {
		return nil, connErr
	}

	sql.Register(
		instrumentedDbDriverName,
		instrumentedsql.WrapDriver(
			connector.Driver(),
			instrumentedsql.WithTracer(instrumentedsqlopentracing.NewTracer()),
			instrumentedsql.WithLogger(
				instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
					logging.SugaredLog.Infof("%s %v", msg, keyvals)
				})),
		),
	)
	db, dbErr := sql.Open(
		instrumentedDbDriverName,
		fmt.Sprintf(dbConnectionStringFormat,
			cfg.dbHost, cfg.dbPort,
			cfg.dbUsername, cfg.dbPassword, cfg.dbName,
			cfg.dbSslMode,
		),
	)
	if dbErr != nil {
		return nil, dbErr
	}

	_, tableErr := db.Exec(createTableQuery)
	if tableErr != nil {
		return nil, tableErr
	}

	return db, nil
}

// WARN: does not work
// imports
// 		"contrib.go.opencensus.io/integrations/ocsql"
// func NewWithOcsqlTracing() (*sql.DB, error) {
// 	logging.Log.Info("Create new DB connector with tracing")
//
// 	cfg := loadConfig()
//
// 	var connector driver.Connector
// 	var connErr error
//
// 	// Get a database driver.Connector for a fixed configuration.
// 	connector, connErr = pq.NewConnector(fmt.Sprintf(dbConnectionString,
// 		cfg.dbUsername, cfg.dbPassword,
// 		cfg.dbHost, cfg.dbPort,
// 		cfg.dbName, cfg.dbSslMode,
// 	))
// 	if connErr != nil {
// 		return nil, connErr
// 	}
//
// 	// Wrap the driver.Connector with ocsql.
// 	connector = ocsql.WrapConnector(connector, ocsql.WithAllTraceOptions())
//
// 	// Use the wrapped driver.Connector.
// 	db := sql.OpenDB(connector)
//
// 	_, tableErr := db.Exec(createTableQuery)
// 	if tableErr != nil {
// 		return nil, tableErr
// 	}
//
// 	return db, nil
// }
