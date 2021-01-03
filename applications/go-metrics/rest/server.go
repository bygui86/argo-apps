package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/go-metrics/monitoring"
	"github.com/bygui86/go-metrics/logging"

	"github.com/gorilla/mux"
)

// RestServer -
type RestServer struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
}

// NewRestServer - Create new REST server
func NewRestServer(customMetrics monitoring.ICustomMetrics) (*RestServer, error) {

	logging.Log.Infoln("[REST] Setup new REST server...")

	// create config
	cfg, err := newConfig()
	if err != nil {
		return nil, err
	}

	// create custom metrics
	addCustomMetrics(customMetrics)

	// create router
	router := newRouter(cfg, customMetrics)

	// create http server
	httpServer := newHTTPServer(cfg.RestHost, cfg.RestPort, router)

	return &RestServer{
		Config:     cfg,
		Router:     router,
		HTTPServer: httpServer,
	}, nil
}

// newRouter -
func newRouter(cfg *Config, customMetrics monitoring.ICustomMetrics) *mux.Router {

	logging.Log.Debugln("[REST] Setup new Router config...")

	router := mux.NewRouter().StrictSlash(false)

	router.
		HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
			echo(w, r, customMetrics)
		}).
		Methods(http.MethodGet)
	router.
		HandleFunc("/echo/{msg}", func(w http.ResponseWriter, r *http.Request) {
			echo(w, r, customMetrics)
		}).
		Methods(http.MethodGet)

	return router
}

// newHttpServer -
func newHTTPServer(host string, port int, router *mux.Router) *http.Server {

	logging.Log.Debugf("[REST] Setup new HTTP server on port %d...", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

// Start - Start REST server
func (s *RestServer) Start() {

	logging.Log.Infoln("[REST] Start REST server...")

	// TODO add a channel to communicate if everything is right
	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			logging.Log.Errorln("[REST] Error starting REST server:", err)
		}
	}()

	logging.Log.Infoln("[REST] REST server listen on port", s.Config.RestPort)
}

// Shutdown - Shutdown REST server
func (s *RestServer) Shutdown() {

	logging.Log.Warnln("[REST] Shutdown REST server...")
	if s.HTTPServer != nil {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		s.HTTPServer.Shutdown(ctx)
	}
}
