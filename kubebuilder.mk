# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

DOCKER ?= docker
# Image to build protobugs
PROTO_IMG ?= k8s.gcr.io/kube-cross:v1.13.6-1

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# TOP is the current directory where this Makefile lives.
TOP := $(dir $(firstword $(MAKEFILE_LIST)))
# ROOT is the root of the mkdocs tree.
ROOT := $(abspath $(TOP))

# enable Go modules
export GO111MODULE=on

CONTROLLER_GEN=GOFLAGS=-mod=vendor go run sigs.k8s.io/controller-tools/cmd/controller-gen

all: manager

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: update
update: fmt vet generate proto

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate:
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

# Generate protobufs
.PHONY: proto
proto:
	$(DOCKER) run -it \
		--mount type=bind,source=$(ROOT),target=/go/src/sigs.k8s.io/service-apis  \
		--env GOPATH=/go \
		--env GOCACHE=/go/.cache \
		--rm \
		--user "$(shell id -u):$(shell id -g)" \
		-w /go/src/sigs.k8s.io/service-apis \
		$(PROTO_IMG) \
		hack/update-proto.sh

# Verify protobuf generation
.PHONY: verify-proto
verify-proto:
	$(DOCKER) run \
		--mount type=bind,source=$(ROOT),target=/realgo/src/sigs.k8s.io/service-apis \
		--env GOPATH=/go \
		--env GOCACHE=/go/.cache \
		--rm \
		--user "$(shell id -u):$(shell id -g)" \
		-w /go \
		$(PROTO_IMG) \
		/bin/bash -c "mkdir -p src/sigs.k8s.io/service-apis && \
			cp -r /realgo/src/sigs.k8s.io/service-apis/ src/sigs.k8s.io && \
			cd src/sigs.k8s.io/service-apis && \
			hack/update-proto.sh && \
			diff -r api /realgo/src/sigs.k8s.io/service-apis/api"


# Build the docker image
docker-build: test
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}
