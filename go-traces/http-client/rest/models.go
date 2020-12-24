package rest

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Server struct {
	config     *config
	router     *mux.Router
	httpServer *http.Server
	baseURL    *url.URL
	restClient *http.Client
	running    bool
}

type config struct {
	restServerHost string
	restServerPort int
	restHost       string
	restPort       int
}
