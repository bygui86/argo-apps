package grpc_interface

import (
	"fmt"
	"net"
	"time"

	opentracinggrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/bygui86/go-traces/grpc-server/logging"
)

const (
	grpcListenerNetwork = "tcp"
)

func New() (*Server, error) {
	logging.Log.Info("Create new gRPC server")

	cfg := loadConfig()

	listener, listErr := net.Listen(grpcListenerNetwork, cfg.address)
	if listErr != nil {
		return nil, listErr
	}

	var grpcServer *grpc.Server
	tracer := opentracing.GlobalTracer()
	if tracer != nil {
		// do not log spans
		// grpcServer = grpc.NewServer(
		// 	grpc.UnaryInterceptor(opentracinggrpc.OpenTracingServerInterceptor(tracer)),
		// 	grpc.StreamInterceptor(opentracinggrpc.OpenTracingStreamServerInterceptor(tracer)),
		// )

		// log spans
		grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(opentracinggrpc.OpenTracingServerInterceptor(tracer, opentracinggrpc.LogPayloads())),
			grpc.StreamInterceptor(opentracinggrpc.OpenTracingStreamServerInterceptor(tracer, opentracinggrpc.LogPayloads())),
		)
	} else {
		grpcServer = grpc.NewServer()
	}

	server := &Server{
		config:     cfg,
		listener:   listener,
		grpcServer: grpcServer,
		running:    false,
	}

	RegisterHelloServiceServer(grpcServer, server)
	return server, nil
}

func (s *Server) Start() error {
	logging.Log.Info("Start gRPC server")

	if s.listener != nil && s.grpcServer != nil && !s.running {
		var err error
		go func() {
			err = s.grpcServer.Serve(s.listener)
			if err != nil {
				logging.SugaredLog.Errorf("REST server start failed: %s", err.Error())
			}
		}()
		if err != nil {
			return err
		}
		s.running = true
		logging.SugaredLog.Infof("gRPC server started")
		return nil
	}

	return fmt.Errorf("gRPC server start failed: gRPC listener not initialized, gRPC server not initialized or already running")
}

func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown gRPC server, timeout %d", timeout)

	if s.listener != nil && s.grpcServer != nil && s.running {
		s.grpcServer.Stop()

		time.Sleep(time.Duration(timeout) * time.Second)

		err := s.listener.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down gRPC listener: %s", err.Error())
		}

		s.running = false
		return
	}

	logging.Log.Error("gRPC server shutdown failed: gRPC listener not initialized, gRPC server not initialized or not running")
}
