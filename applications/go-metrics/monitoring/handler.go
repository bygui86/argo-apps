package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// getMetricsHandler - Return an HTTP handler for the moniroting
func getMetricsHandler() http.Handler {

	return promhttp.Handler()
}
