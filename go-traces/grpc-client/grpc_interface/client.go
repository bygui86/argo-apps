package grpc_interface

import (
	"context"
	"fmt"
	"time"

	opentracinggrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/bygui86/go-traces/grpc-client/logging"
)

func New() (*Client, error) {
	logging.Log.Info("Create new gRPC client")

	cfg := loadConfig()

	tracer := opentracing.GlobalTracer()
	var connection *grpc.ClientConn
	var err error
	if tracer != nil {
		// do not log spans
		// connection, err = grpc.Dial(
		// 	cfg.grpcServerAddress,
		// 	grpc.WithInsecure(),
		// 	grpc.WithUnaryInterceptor(opentracinggrpc.OpenTracingClientInterceptor(tracer)),
		// 	grpc.WithStreamInterceptor(opentracinggrpc.OpenTracingStreamClientInterceptor(tracer)),
		// )

		// log spans
		connection, err = grpc.Dial(
			cfg.grpcServerAddress,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(opentracinggrpc.OpenTracingClientInterceptor(tracer, opentracinggrpc.LogPayloads())),
			grpc.WithStreamInterceptor(opentracinggrpc.OpenTracingStreamClientInterceptor(tracer, opentracinggrpc.LogPayloads())),
		)
	} else {
		connection, err = grpc.Dial(
			cfg.grpcServerAddress,
			grpc.WithInsecure(),
		)
	}
	if err != nil {
		return nil, err
	}

	logging.SugaredLog.Debugf("State: %v", connection.GetState())
	logging.SugaredLog.Debugf("Target: %v", connection.Target())

	return &Client{
		config:         cfg,
		ctx:            context.Background(),
		grpcConnection: connection,
		running:        false,
	}, nil
}

func (c *Client) Start() error {
	logging.Log.Info("Start gRPC client")

	if c.grpcConnection != nil && !c.running {
		c.helloServiceClient = NewHelloServiceClient(c.grpcConnection)
		go c.startGreeting()
		c.running = true
		logging.SugaredLog.Infof("gRPC server started")
		return nil
	}

	return fmt.Errorf("gRPC client start failed: gRPC connection not initialized or already running")
}

func (c *Client) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown gRPC client, timeout %d", timeout)

	if c.grpcConnection != nil && c.ticker != nil && c.running {
		c.ticker.Stop()
		c.ctx.Done()

		time.Sleep(time.Duration(timeout) * time.Second)

		err := c.grpcConnection.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down gRPC client: %s", err.Error())
		}

		c.running = false
		return
	}

	logging.Log.Error("gRPC client shutdown failed: gRPC connection not initialized or not running")
}
