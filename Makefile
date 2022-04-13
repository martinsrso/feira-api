GO ?= go
GOGENERATE ?= $(GO) generate
GOINSTALL ?= $(GO) install
GOBUILD ?= $(GO) build
GOMOD ?= $(GO) mod
GORUN ?= $(GO) run
GOTEST ?= $(GO) test
GOTOOL ?= $(GO) tool

GIT ?= git
GITDIFF ?= $(GIT) diff

################################################################################
## Main make targets
################################################################################
.PHONY: run
run:
	$(GORUN) main.go

################################################################################
## Docker make targets
################################################################################
.PHONY: docker/build
docker/build:
	docker build --rm -t feira-api .

.PHONY: docker/run
docker/run: docker/build
	docker run --rm -dp 8888:8888 feira-api

.PHONY: docker/up
docker/up:
	docker-compose build
	docker-compose up -d

.PHONY: docker/down
docker/down:
	docker-compose down

.PHONY: docker/clean
docker/clean:
	docker stop `docker ps -qa`
	docker rm `docker ps -qa`
	docker rmi -f `docker images -qa `
	docker network rm `docker network ls -q`

################################################################################
## Go-like targets
################################################################################
.PHONY: build
build:
	$(GOBUILD) -mod vendor -a -installsuffix cgo -o . .

.PHONY: test
test:
	$(GOTEST) -failfast -coverprofile=coverage.out ./... $(SILENT_CMD_SUFFIX)

.PHONY: test/race
test/race:
	$(GOTEST) -race ./... $(SILENT_CMD_SUFFIX)

.PHONY: cover
cover: cover/text

.PHONY: cover/html
cover/html:
	$(GOTOOL) cover -html=coverage.out

.PHONY: cover/text
cover/text:
	$(GOTOOL) cover -func=coverage.out

.PHONY: mocks
mocks:
	$(GORUN) github.com/golang/mock/mockgen -source=domain/market.go -destination=domain/mocks/mock_market.go -package mock

.PHONY: vendors
vendors:
	$(GOMOD) vendor
	$(GOMOD) tidy

.PHONY: git/diff
git/diff:
	@if ! $(GITDIFF) --quiet; then \
		printf 'Found changes on local workspace. Please run this target and commit the changes\n' ; \
		exit 1; \
	fi
