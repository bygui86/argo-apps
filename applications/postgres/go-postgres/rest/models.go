package rest

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type config struct {
	DbHost     string
	DbPort     int
	DbUsername string
	DbPassword string
	DbName     string
	DbSslMode  string
	RestHost   string
	RestPort   int
}

type Server struct {
	config       *config
	Router       *mux.Router
	httpServer   *http.Server
	DbConnection *sql.DB
	running      bool
}
