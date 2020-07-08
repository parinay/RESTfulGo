PROJECT_ROOT := github.com/parinay/RESTfulGo
BUILD_PATH := bin
DOCKERFILE_PATH := $(CURDIR)/docker
DOCKERFILE   := $(DOCKERFILE_PATH)/Dockerfile
#
GIT_COMMIT		:= $(shell git describe --tags --dirty=-unsupported --always || echo pre-commit)
IMAGE_VERSION       ?= $(GIT_COMMIT)
IMAGE 			:= restful.go.api
#
CLIENT_DOCKERFILE   := $(DOCKERFILE_PATH)/Dockerfile
#
BIN := crud
# configuration for building on host machine
GO_CACHE		:= -pkgdir $(BUILD_PATH)/go-cache
GO_BUILD_FLAGS		?= $(GO_CACHE) -i -v -race
GO_TEST_FLAGS		?= -v -cover -race
GO_PACKAGES		:= $(shell go list ./... | grep -Ev 'vendor|testclient')
GO_FILES                := $(shell find . -type f -name '*.go' ! -path "./vendor/*" ! -path "./pkg/pb/*")

.PHONY: all
all: docker

.PHONY: imports
imports:
	@goimports -w $(GO_FILES)

.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)
.PHONY: lint
lint:
	@! gofmt -l . | grep -v vendor/

.PHONY: build
build:
	@go build $(GO_BUILD_FLAGS) -v -o $(BUILD_PATH)/$(BIN) crud/*.go

.PHONY: docker
docker:
	@docker build --build-arg VERSION=$(GIT_COMMIT) -f $(DOCKERFILE) -t $(IMAGE):$(IMAGE_VERSION) .
	@docker image prune -f  || true
.PHONY: clean
clean:
	$(info Removing Docker image: $(CLIENT_IMAGE):$(IMAGE_VERSION))
	@docker rmi -f $(shell docker images -q $(IMAGE):$(IMAGE_VERSION)) 2>/dev/null || true

	rm -f  $(BUILD_PATH)/$(BIN)
