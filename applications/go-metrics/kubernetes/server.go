package kubernetes

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/go-metrics/logging"

	"github.com/gorilla/mux"
)

// KubeServer -
type KubeServer struct {
	Config     *Config
	Router     *mux.Router
	HttpServer *http.Server
}

// NewKubeServer - Create new Kubernetes REST server
func NewKubeServer() (*KubeServer, error) {

	logging.Log.Infoln("[KUBERNETES] Setup new REST server...")

	// create config
	cfg, err := newConfig()
	if err != nil {
		return nil, err
	}

	// create router
	router := newRouter()

	// create http server
	httpServer := newHttpServer(cfg.RestHost, cfg.RestPort, router)

	return &KubeServer{
		Config:     cfg,
		Router:     router,
		HttpServer: httpServer,
	}, nil
}

// newRouter -
func newRouter() *mux.Router {

	logging.Log.Debugln("[KUBERNETES] Setup new Router config...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/live", livenessHandler)
	router.HandleFunc("/ready", readynessHandler)
	return router
}

// newHttpServer -
func newHttpServer(host string, port int, router *mux.Router) *http.Server {

	logging.Log.Debugf("[KUBERNETES] Setup new HTTP server on port %d...", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

// Start - Start Kubernetes REST server
func (s *KubeServer) Start() {

	logging.Log.Infoln("[KUBERNETES] Start REST server...")

	// TODO add a channel to communicate if everything is right
	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil {
			logging.Log.Errorln("[KUBERNETES] Error starting REST server:", err)
		}
	}()

	logging.Log.Infoln("[KUBERNETES] REST server listen on port", s.Config.RestPort)
}

// Shutdown - Shutdown Kubernetes REST server
func (s *KubeServer) Shutdown() {

	logging.Log.Warnln("[KUBERNETES] Shutdown REST server...")
	if s.HttpServer != nil {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		s.HttpServer.Shutdown(ctx)
	}
}
