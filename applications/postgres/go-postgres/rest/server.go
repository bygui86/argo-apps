package rest

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/bygui86/go-postgres/database"
	"github.com/bygui86/go-postgres/logging"
)

const (
	dbConnectionStringFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dbDriverName             = "postgres"

	httpServerHostFormat          = "%s:%d"
	httpServerWriteTimeoutDefault = time.Second * 15
	httpServerReadTimeoutDefault  = time.Second * 15
	httpServerIdelTimeoutDefault  = time.Second * 60
)

func NewServer() (*Server, error) {
	logging.Log.Debug("Create new REST server")

	cfg := loadConfig()

	dbConnection, dbErr := sql.Open(
		dbDriverName,
		fmt.Sprintf(dbConnectionStringFormat,
			cfg.DbHost, cfg.DbPort,
			cfg.DbUsername, cfg.DbPassword, cfg.DbName,
			cfg.DbSslMode,
		),
	)
	if dbErr != nil {
		return nil, dbErr
	}

	_, tableErr := dbConnection.Exec(database.CreateTableQuery)
	if tableErr != nil {
		return nil, tableErr
	}

	server := &Server{
		config:       cfg,
		DbConnection: dbConnection,
	}

	server.setupRouter()
	server.setupHTTPServer()
	return server, nil
}

func (s *Server) Start() {
	logging.Log.Info("Start REST server")

	if s.httpServer != nil && !s.running {
		go func() {
			err := s.httpServer.ListenAndServe()
			if err != nil {
				logging.SugaredLog.Errorf("Error starting REST server: %s", err.Error())
			}
		}()
		s.running = true
		logging.SugaredLog.Infof("REST server listening on port %d", s.config.RestPort)
		return
	}

	logging.Log.Error("REST server start failed: HTTP server not initialized or HTTP server already running")
}

func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown REST server, timeout %d", timeout)

	if s.httpServer != nil && s.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down REST server: %s", err.Error())
		}
		s.running = false
		return
	}

	logging.Log.Error("REST server shutdown failed: HTTP server not initialized or HTTP server not running")
}
