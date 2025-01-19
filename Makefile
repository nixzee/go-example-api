# Constants
CONTAINER_REGISTRY="some_registry"# TODO: Point to your registry
IMAGE_NAME="go-example-api"
IMAGE_VERSION="v1.0.0"
FULL_IMAGE="$(CONTAINER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_VERSION)"

# Variables
SHA=""
SHORT_SHA=""
OS_INFO=""
DOCKER_VERSION=""
GOLANG_VERSION=""

# Check if in GitHub
ACTION=false
ifneq ($(GITHUB_RUN_NUMBER),)
	ACTION=true
endif

# Get the SHA
ifeq ($(ACTION), true)
	SHA=$(GITHUB_SHA)
	SHORT_SHA=$(shell git rev-parse --short=4 $(GITHUB_SHA))
else
	SHA=$(shell git log -1 --format=%H)
	SHORT_SHA=$(shell git log -1 --format=%h)
endif

# Get Version info
OS_INFO=$(shell uname -a)
DOCKER_VERSION=$(shell docker --version 2>/dev/null | cut -d " " -f 3 | cut -d "," -f 1)
GOLANG_VERSION=$(shell go version 2>/dev/null | cut -d " " -f 3)

.PHONY: about
about:
	@echo "[GIT]"
	@echo "SHA: $(SHA)"
	@echo "SHORT SHA: $(SHORT_SHA)"
	@echo ""
	@echo "[ENVIRONMENT]"
	@echo "OS INFO: $(OS_INFO)"
	@echo "DOCKER VERSION: $(DOCKER_VERSION)"
	@echo "GOLANG VERSION: $(GOLANG_VERSION)"
	@echo ""
	@echo "[IMAGE]"
	@echo "IMAGE FULL: $(FULL_IMAGE)"

.PHONY: build_amd64
build_amd64:
	@echo "Building $(FULL_IMAGE)_AMD64"
	@docker build -f ./docker/dockerfile.build -t $(FULL_IMAGE)_AMD64 \
		--progress=plain \
		--build-arg GOARCH="amd64" \
		--build-arg APP_GIT_COMMIT="$(SHORT_SHA)" \
		--build-arg APP_VERSION="$(IMAGE_VERSION)" \
		.

.PHONY: build_arm64
build_arm64:
	@echo "Building $(FULL_IMAGE)_ARM64"
	@docker build -f ./docker/dockerfile.build -t $(FULL_IMAGE)_ARM64 \
		--progress=plain \
		--build-arg GOARCH="arm64" \
		--build-arg APP_GIT_COMMIT="$(SHORT_SHA)" \
		--build-arg APP_VERSION="$(IMAGE_VERSION)" \
		.

.PHONY: push_amd64
push_amd64:
	@echo "Pushing $(FULL_IMAGE)_AMD64"
	@docker push $(FULL_IMAGE)_AMD64

.PHONY: push_arm64
push_arm64:
	@echo "Pushing $(FULL_IMAGE)_ARM64"
	@docker push $(FULL_IMAGE)_ARM64

.PHONY: create_manifest
create_manifest:
	@echo "Creating manifest"
	@docker manifest create $(FULL_IMAGE) \
		$(FULL_IMAGE)_AMD64 --os linux --arch amd64 \
		$(FULL_IMAGE)_ARM64 --os linux --arch arm64

.PHONY: push_manifest
push_manifest:
	@echo "Pushing manifest"
	@docker manifest push $(FULL_IMAGE)

.PHONY: run_docker
run_docker:
	@echo "Running container with $(FULL_IMAGE)_AMD64"
	@docker run -it --rm \
		-e ACCOUNT_NAME=$(ACCOUNT_NAME) \
		-e SAS_TOKEN=$(SAS_TOKEN) \
		-p 8080:8080 \
		$(FULL_IMAGE)