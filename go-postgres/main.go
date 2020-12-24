package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-postgres/logging"
	"github.com/bygui86/go-postgres/rest"
)

func main() {
	logging.Log.Info("Start go-traces")

	restServer := startRestServer()

	logging.Log.Info("go-traces up&running")

	startSysCallChannel()

	shutdownAndWait(restServer, 3)
}

func startRestServer() *rest.Server {
	logging.Log.Debug("Start REST server")

	server, err := rest.NewServer()
	if err != nil {
		logging.SugaredLog.Errorf("REST server creation failed: %s", err.Error())
		os.Exit(501)
	}
	logging.Log.Debug("REST server successfully created")

	server.Start()
	logging.Log.Debug("REST server successfully started")

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(restServer *rest.Server, timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)
	restServer.Shutdown(timeout)
	time.Sleep(time.Duration(timeout+1) * time.Second)
}
