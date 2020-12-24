
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

### jaeger

start-jaeger :		## Run Jaeger (all-in-one) in a container
	docker run -d --rm --name jaeger \
		-p 5775:5775/udp \
		-p 5778:5578 \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 9411:9411 \
		-p 14268:14268 \
		-p 14269:14269 \
		-p 16686:16686 \
		jaegertracing/all-in-one

stop-jaeger :		## Stop Jaeger (all-in-one) in container
	docker stop jaeger

open-jaeger-ui :		## Open Jaeger UI in browser
	open http://localhost:16686

### zipkin

start-zipkin :		## Run Zipkin in a container
	docker run -d --rm --name zipkin \
		-p 9411:9411 \
		openzipkin/zipkin

stop-zipkin :		## Stop Zipkin in container
	docker stop zipkin

open-zipkin-ui :		## Open Zipkin UI in browser
	open http://localhost:9411

### postgres

start-postgres :		## Run PostgreSQL in a container
	cd http-server/ && make start-postgres

stop-postgres :		## Stop PostgreSQL in container
	cd http-server/ && make stop-postgres

### kafka

start-kafka :		## Run Apache Zookeeper and Apache Kafka in containers
	cd kafka-producer/ && make start-kafka

stop-kafka :		## Stop Apache Zookeeper and Apache Kafka in containers
	cd kafka-producer/ && make stop-kafka

### kubemq

start-kubemq :		## Run Minikube and deploy KubeMQ
	cd kubemq-producer/ && make start-kubemq

stop-kubemq :		## Stop Minikube and KubeMQ
	cd kubemq-producer/ && make stop-kubemq

proxy-kubemq :		## Proxy KubeMQ
	cd kubemq-producer/ && make proxy-kubemq

open-kubemq-ui :		## Open KubeMQ UI in browser
	cd kubemq-producer/ && make open-kubemq-ui

## applications

### http

start-http-server :		## Run HTTP server
	cd http-server/ && make start

start-http-client :		## Run HTTP client
	cd http-client/ && make start

### kafka

start-kafka-consumer :		## Run Kafka consumer
	cd kafka-consumer/ && make start

start-kafka-producer :		## Run Kafka producer
	cd kafka-producer/ && make start

### kubemq

start-kubemq-consumer :		## Run KubeMQ consumer
	cd kubemq-consumer/ && make start

start-kubemq-producer :		## Run KubeMQ producer
	cd kubemq-producer/ && make start

### grpc

build-protobuf :		## Compile protobuf
	rm -f ./grpc-protobuf/*.pb.go
	rm -f ./grpc-server/grpc_interface/*.pb.go
	rm -f ./grpc-client/grpc_interface/*.pb.go
	protoc --proto_path=./grpc-protobuf/ --go_out=plugins=grpc:./grpc-protobuf ./grpc-protobuf/*
	cp ./grpc-protobuf/*.pb.go ./grpc-server/grpc_interface/
	cp ./grpc-protobuf/*.pb.go ./grpc-client/grpc_interface/

start-grpc-server :		## Run gRPC server
	cd grpc-server/ && make start

start-grpc-client :		## Run gRPC client
	cd grpc-client/ && make start

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
