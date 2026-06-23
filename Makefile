# ----------------------------------------------------------------------------------------------------------------------------------------- Go Configs

.PHONY: test
test:
	go test ./... -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: cover
cover: 
	go test ./... -coverprofile=coverage.out
	@echo "coverage written to coverage.out"

# Define an empty variable for extra args
ARGS=

# ----------------------------------------------------------------------------------------------------------------------------------------- run cli

.PHONY: cli
cli: ## Build the CLI binary.
	@echo "Executing command with: $(ARGS)"
	@echo "---------------------------------------------------------------------------------"
	go run cmd/eactl/main.go $(ARGS)

# ----------------------------------------------------------------------------------------------------------------------------------------- Docker Configs

REGISTRY ?= localhost:5000
REPOSITORY ?= ea
TAG ?= latest

.PHONY: build
build: ## Build docker image with the manager.
	docker build -t ${REPOSITORY}:${TAG} -t ${REGISTRY}/${REPOSITORY}:${TAG} . 

.PHONY: push
push: build ## Push docker image with the manager.
	docker push ${REGISTRY}/${REPOSITORY}:${TAG}
	
.PHONY: run
run: build ## Run docker image with the manager.
	@echo "Executing command with: $(ARGS)"
	@echo "---------------------------------------------------------------------------------"
	docker run -d --rm -p 80:80 ${REPOSITORY}:${TAG} $(ARGS)

