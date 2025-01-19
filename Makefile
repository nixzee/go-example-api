# Constants
CONTAINER_REGISTRY=""
IMAGE_NAME="go-example-api"
IMAGE_VERSION="v1.0.0"

# Variables
SHA=""
SHORT_SHA=""
OS_INFO=""
DOCKER_VERSION=""
GOLANG_VERSION=""

# Check if in GitHub
ACTION=false
ifeq ($(GITHUB_RUN_NUMBER),)
	ACTION=true
endif

# Get the SHA
ifeq ($(ACTION), true)
	SHA=${GITHUB_SHA}
	SHORT_SHA=$(shell git rev-parse --short=4 ${GITHUB_SHA})
else
	SHA=$(shell git log -1 --format=%H)
	SHORT_SHA=$(shell git log -1 --format=:%h)
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

