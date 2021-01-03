
# VARIABLES
ES_SECRET := logging-es-elastic-user


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## deploy

create-pvc :		## Create Kubernetes PVC for Grafana
	kubectl apply -f pvc.yaml

deploy :		## Deploy Grafana
	kustomize build . | kubectl apply -f -
	@echo ""
	@echo "/!\"
	@echo "Remember to manually import dashboards!"
	@echo "/!\"
	@echo ""
	@echo ""

deploy-all : create-pvc deploy		## Deploy Grafana environment


## delete

delete-pvc :		## Delete Kubernetes PVC for Grafana
	kubectl delete pvc -l app=grafana,domain=monitoring,component=grafana -n monitoring

delete :		## Delete Grafana
	kustomize build . | kubectl delete -f -

delete-all : delete		## Delete Grafana environment


## ops

port-forw :		## Open port forwarding to Grafana
	@kubectl port-forward svc/grafana 3000 -n monitoring


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
