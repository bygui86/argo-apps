
# VARIABLES
export GO111MODULE=on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

build :		## Build application
	go build

test : run-postgres		## Run tests (required a running PostgreSQL instance)
	go test ./...

run : run-postgres		## Run application from source code
	godotenv -f local.env go run main.go

run-postgres :		## Run PostgreSQL in a container
	-@docker run -d --name postgres \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		postgres

docker-build :		## Build container
	docker build . -t bygui86/go-postgres:latest

docker-push :		## Push container to Docker hub
	docker push bygui86/go-postgres:latest


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
