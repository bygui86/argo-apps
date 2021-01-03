package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bygui86/go-metrics/kubernetes"
	"github.com/bygui86/go-metrics/logging"
	"github.com/bygui86/go-metrics/monitoring"
	"github.com/bygui86/go-metrics/rest"
)

// main -
func main() {

	logging.Log.Infoln("[MAIN] Starting echo-server...")

	monitorServer := startMonitor()
	defer monitorServer.Shutdown()

	restServer := startRest(monitorServer.CustomMetrics)
	defer restServer.Shutdown()

	kubeServer := startKubernetes()
	defer kubeServer.Shutdown()

	logging.Log.Infoln("[MAIN] echo-server ready!")

	startSysCallChannel()
}

// startMonitor -
func startMonitor() *monitoring.MonitorServer {

	server, err := monitoring.NewMonitorServer()
	if err != nil {
		logging.Log.Errorf("[MAIN] Monitoring server creation failed: %s", err.Error())
		os.Exit(404)
	}
	logging.Log.Debugln("[MAIN] Monitoring server successfully created")

	server.Start()
	logging.Log.Debugln("[MAIN] Monitoring successfully started")

	return server
}

// startRest -
func startRest(customMetrics monitoring.ICustomMetrics) *rest.RestServer {

	server, err := rest.NewRestServer(customMetrics)
	if err != nil {
		logging.Log.Errorf("[MAIN] Echo server creation failed: %s", err.Error())
		os.Exit(404)
	}
	logging.Log.Debugln("[MAIN] Echo server successfully created")

	server.Start()
	logging.Log.Debugln("[MAIN] Echo successfully started")

	return server
}

// startKubernetes -
func startKubernetes() *kubernetes.KubeServer {

	server, err := kubernetes.NewKubeServer()
	if err != nil {
		logging.Log.Errorf("[MAIN] Kubernetes server creation failed: %s", err.Error())
		os.Exit(404)
	}
	logging.Log.Debugln("[MAIN] Kubernetes server successfully created")

	server.Start()
	logging.Log.Debugln("[MAIN] Kubernetes successfully started")

	return server
}

// startSysCallChannel -
func startSysCallChannel() {

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logging.Log.Warnln("[MAIN] Termination signal received!")
}
