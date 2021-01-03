
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## deploy

deploy-operator :		## Deploy Prometheus Operator
	kustomize build 20_operator/ | kubectl apply -f -

deploy-kube-state :		## Deploy Kube-State-Metrics
	kustomize build 30_kube-state-metrics/ | kubectl apply -f -

deploy-node-exp :		## Deploy Node-Exporter
	kustomize build 40_node-exporter/ | kubectl apply -f -

deploy-prom-adapt :		## Deploy Prometheus-Adapter
	kustomize build 50_prometheus-adapter/ | kubectl apply -f -

deploy-prom-config :		## Deploy Prometheus configurations for all namespaces
	kustomize build 60_specific-namespaces/ | kubectl apply -f -
	kubectl apply -f 61_additionals/

deploy-rules :		## Deploy Prometheus (alerting) rules
	kustomize build 62_rules | kubectl apply -f -

deploy-prom :		## Deploy Prometheus
	kustomize build 70_prometheus/ | kubectl apply -f -

deploy-all : deploy-operator deploy-kube-state deploy-node-exp deploy-prom-adapt deploy-prom-config deploy-rules deploy-prom		## Deploy only Prometheus environment

## delete

delete-operator :		## Delete Prometheus Operator
	kustomize build 20_operator/ | kubectl delete -f -

delete-prom-adapt :		## Delete Prometheus-Adapter
	kustomize build 50_prometheus-adapter/ | kubectl delete -f -

delete-node-exp :		## Delete Node-Exporter
	kustomize build 40_node-exporter/ | kubectl delete -f -

delete-kube-state :		## Delete Kube-State-Metrics
	kustomize build 30_kube-state-metrics/ | kubectl delete -f -

delete-prom-config :		## Delete Prometheus configurations for all namespaces
	kustomize build 60_specific-namespaces/ | kubectl delete -f -
	kubectl delete -f 61_additionals/

delete-rules :		## Delete Prometheus (alerting) rules
	kustomize build 62_rules | kubectl delete -f -

delete-prom :		## Delete Prometheus
	kustomize build 70_prometheus/ | kubectl delete -f -

delete-all : delete-prom delete-rules delete-prom-config delete-kube-state delete-node-exp delete-prom-adapt delete-operator		## Delete whole Prometheus environment


## pvc

delete-pvc :		## Delete Kubernetes PersistentVolumeClaim for Prometheus environment
	kubectl delete pvc -l app=prometheus


## ops

port-forw-prom :		## Open port forwarding to Prometheus
	@kubectl port-forward svc/prometheus 9090 -n monitoring

# port-forw-pushgateway :		## Open port forwarding to Prometheus pushgateway
# 	@kubectl port-forward svc/prometheus-pushgateway 9091 -n monitoring


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
	@echo ""
