package kubernetes

import (
	"github.com/bygui86/go-metrics/envvars"
	"github.com/bygui86/go-metrics/logging"
)

const (
	// Environment variables -
	kubeHostEnvVar            = "KUBE_HOST"
	kubePortEnvVar            = "KUBE_PORT"
	kubeShutdownTimeoutEnvVar = "KUBE_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	kubeHostDefault     = "localhost"
	kubePortDefault     = 9091
	kubeShutdownTimeout = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logging.Log.Debugln("[KUBERNETES] Setup new Kubernetes config...")

	return &Config{
		RestHost:        envvars.GetStringEnv(kubeHostEnvVar, kubeHostDefault),
		RestPort:        envvars.GetIntEnv(kubePortEnvVar, kubePortDefault),
		ShutdownTimeout: envvars.GetIntEnv(kubeShutdownTimeoutEnvVar, kubeShutdownTimeout),
	}, nil
}
