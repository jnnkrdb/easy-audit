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

# ----------------------------------------------------------------------------------------------------------------------------------------- run cli
# Define an empty variable for extra args
ARGS=
.PHONY: cli
cli: ## Build the CLI binary.
	@echo "Executing command with: $(ARGS)"
	@echo "---------------------------------------------------------------------------------"
	go run cmd/cli/main.go $(ARGS)

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
	docker run -d --rm -p 80:80 ${REPOSITORY}:${TAG} easy-audit --log-level=debug --verbose --log-format=text --database-dsn="file:/data/audits.db" --database-driver=sqlite

