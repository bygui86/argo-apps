package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-postgres/logging"
)

func (s *Server) setupRouter() {
	logging.Log.Debug("Create new router")

	s.Router = mux.NewRouter().StrictSlash(true)

	s.Router.HandleFunc("/products", s.getProducts).Methods(http.MethodGet)
	s.Router.HandleFunc("/products/{id:[0-9]+}", s.getProduct).Methods(http.MethodGet)
	s.Router.HandleFunc("/products", s.createProduct).Methods(http.MethodPost)
	s.Router.HandleFunc("/products/{id:[0-9]+}", s.updateProduct).Methods(http.MethodPut)
	s.Router.HandleFunc("/products/{id:[0-9]+}", s.deleteProduct).Methods(http.MethodDelete)
}

func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.config.RestPort)

	if s.config != nil {
		s.httpServer = &http.Server{
			Addr:    fmt.Sprintf(httpServerHostFormat, s.config.RestHost, s.config.RestPort),
			Handler: s.Router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: httpServerWriteTimeoutDefault,
			ReadTimeout:  httpServerReadTimeoutDefault,
			IdleTimeout:  httpServerIdelTimeoutDefault,
		}
		return
	}

	logging.Log.Error("HTTP server creation failed: REST server configurations not initialized")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "rest/json")
	w.WriteHeader(code)
	w.Write(response)
}
