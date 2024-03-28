
# Image URL to use all building/pushing image targets
TAG ?= $(shell hack/lib/flags.sh --version)
IMG ?= webgamedevelop/webgame-api:$(TAG)
LDFLAGS ?= $(shell hack/lib/flags.sh --ldflags) -X 'k8s.io/component-base/version/verflag.programName=webgame-api'

ENVTEST_K8S_VERSION = 1.28.0

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# CONTAINER_TOOL defines the container tool to be used for building images.
# Be aware that the target commands are only tested with Docker which is
# scaffolded by default. However, you might want to replace it to use other
# tools. (i.e. podman)
CONTAINER_TOOL ?= docker

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: swagger
swagger: swag ## Generate swagger file and format swag comments.
	$(SWAG) fmt --dir internal/handlers/api --generalInfo router.go
	$(SWAG) init --parseDependency --dir internal/handlers/api --output internal/handlers/docs --generalInfo router.go

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" go test ./... -coverprofile cover.out

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.54.2
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter & yamllint.
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes.
	$(GOLANGCI_LINT) run --fix

##@ Build

.PHONY: build
build: fmt swagger vet ## Build webgame-api binary.
	go build --ldflags "$(LDFLAGS)" -o $(LOCALBIN)/webgame-api cmd/main.go

.PHONY: run
run: fmt swagger vet ## Run a webgame-api from your host.
	go run ./cmd/main.go \
	    --database-address localhost \
	    --database-password 123456 \
	    --gin-mode debug \
	    --v 2 \
	    --logger-klog-v 0

.PHONY: run-bin
run-bin: build
	$(LOCALBIN)/webgame-api \
	    --database-address localhost \
	    --database-password 123456 \
	    --gin-mode debug \
	    --v 2 \
	    --logger-klog-v 0

.PHONY: import-data
import-data:
	$(LOCALBIN)/webgame-api \
	    --import-initialization-data \
	    --gorm-debug-log-level \
	    --database-address localhost \
	    --database-password 123456 \
	    --v 2 \
	    --logger-klog-v 2

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: fmt swagger vet ## Build docker image with the webgame-api.
	$(CONTAINER_TOOL) build --build-arg LDFLAGS="-s -w $(LDFLAGS)" -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the webgame-api.
	$(CONTAINER_TOOL) push ${IMG}

# PLATFORMS defines the target platforms for the manager image be built to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - be able to use docker buildx. More info: https://docs.docker.com/build/buildx/
# - have enabled BuildKit. More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image to your registry (i.e. if you do not set a valid value via IMG=<myregistry/image:<tag>> then the export will fail)
# To adequately provide solutions that are compatible with multiple platforms, you should consider using this option.
PLATFORMS ?= linux/arm64,linux/amd64
.PHONY: docker-buildx
docker-buildx: fmt swagger vet ## Build and push docker image for the webgame-api for cross-platform support.
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	- $(CONTAINER_TOOL) buildx create --name project-v3-builder
	$(CONTAINER_TOOL) buildx use project-v3-builder
	- $(CONTAINER_TOOL) buildx build --build-arg LDFLAGS="-s -w $(LDFLAGS)" --push --platform=$(PLATFORMS) --tag ${IMG} -f Dockerfile.cross .
	- $(CONTAINER_TOOL) buildx rm project-v3-builder
	 rm Dockerfile.cross

##@ Deployment

.PHONY: deploy
deploy: helm ## Deploy webgame-api component by helm to the K8s cluster specified in ~/.kube/config.
	$(HELM) -n webgame-system upgrade --install --create-namespace webgame-api helm \
	    --set image.image="$(IMG)" \
	    --set log.ginMode="debug" \
	    --set log.levelEnabler=0 \
	    --set log.level=2 \
	    --set log.inspectLevel=2 \
	    --set database.address="mysql.mysql" \
	    --set database.port="3306" \
	    --set database.user="root" \
	    --set database.password="abc123"

.PHONY: undeploy
undeploy: helm ## Undeploy webgame-api component by helm from the K8s cluster specified in ~/.kube/config.
	$(HELM) -n webgame-system uninstall webgame-api

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

## Tool Binaries
ENVTEST ?= $(LOCALBIN)/setup-envtest
SWAG ?= $(LOCALBIN)/swag
HELM ?= $(LOCALBIN)/helm

## Tool Versions
SWAG_VERSION ?= v1.16.2
HELM_VERSION ?= v3.14.2

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	test -s $(ENVTEST) || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest

.PHONY: swag
swag: $(SWAG) ## Download swag locally if necessary.
$(SWAG): $(LOCALBIN)
	test -s $(SWAG) || GOBIN=$(LOCALBIN) go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

.PHONY: helm
helm: $(HELM) ## Download helm locally if necessary.
$(HELM): $(LOCALBIN)
	test -s $(HELM) || GOBIN=$(LOCALBIN) go install helm.sh/helm/v3/cmd/helm@$(HELM_VERSION)
