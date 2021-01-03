package rest

import (
	"net/http"

	"github.com/bygui86/go-metrics/monitoring"
	"github.com/bygui86/go-metrics/logging"

	"github.com/gorilla/mux"
)

const (
	defaultMsg = "Hello world!"
)

// echo -
func echo(w http.ResponseWriter, r *http.Request, customMetrics monitoring.ICustomMetrics) {

	vars := mux.Vars(r)
	msg := vars["msg"]

	if len(msg) == 0 {
		logging.Log.Infof("[REST] Echo of default msg '%s'", defaultMsg)
		w.Write([]byte(defaultMsg))
	} else {
		logging.Log.Infof("[REST] Echo of msg '%s'", msg)
		w.Write([]byte(msg))
	}

	go updateCustomMetrics(customMetrics)
}

// updateCustomMetrics -
func updateCustomMetrics(customMetrics monitoring.ICustomMetrics) {

	if customMetrics != nil {
		customMetrics.IncreaseCounter(opsProcessedKey)
	}
}
