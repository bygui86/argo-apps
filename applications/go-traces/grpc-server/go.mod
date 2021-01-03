module github.com/bygui86/go-traces/grpc-server

go 1.14

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/opentracing-contrib/go-grpc v0.0.0-20191001143057-db30781987df
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/prometheus/client_golang v1.6.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.29.1
)
