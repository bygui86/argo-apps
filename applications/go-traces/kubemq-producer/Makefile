
# VARIABLES
# -


# ENVIRONMENT VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

start-kubemq :		## Run Minikube and deploy KubeMQ
	minikube start --vm-driver=hyperkit --cpus=4 --memory=8192 --disk-size=20g --enable-default-cni --network-plugin=cni --extra-config=kubelet.network-plugin=cni
	@sleep 30s
	kubectl apply -f https://get.kubemq.io/deploy
	@sleep 30s
	kubectl get kubemqclusters,kubemqdashboards -n kubemq

stop-kubemq :		## Stop Minikube and KubeMQ
	minikube stop

proxy-kubemq :		## Proxy KubeMQ
	kubemqctl set cluster proxy

open-kubemq-ui :		## Open KubeMQ UI in browser
	kubemqctl get dashboard

## application

build :		## Build application
	go build

start :		## Run application from source code
	godotenv -f local.env go run main.go

start-binary : build		## Run application from binary
	source local-binary.env && ./kubemq-producer

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
