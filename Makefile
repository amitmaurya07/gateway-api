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

# Enable Go modules.
export GO111MODULE=on

DOCKER ?= docker
# TOP is the current directory where this Makefile lives.
TOP := $(dir $(firstword $(MAKEFILE_LIST)))
# ROOT is the root of the mkdocs tree.
ROOT := $(abspath $(TOP))

all: generate vet fmt verify test

# Run generators for protos, Deepcopy funcs, CRDs, and docs.
.PHONY: generate
generate:
	hack/update-codegen.sh

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Run go test against code
test:
	go test ./...

# Install CRD's and example resources to a pre-existing cluster.
.PHONY: install
install: crd example

# Install the CRD's to a pre-existing cluster.
.PHONY: crd
crd:
	kubectl kustomize config/crd | kubectl apply -f -

# Install the example resources to a pre-existing cluster.
.PHONY: example
example:
	hack/install-examples.sh

# Remove installed CRD's and CR's.
.PHONY: uninstall
uninstall:
	hack/delete-crds.sh

# Run static analysis.
.PHONY: verify
verify:
	hack/verify-all.sh

# Build the documentation.
.PHONY: docs
docs:
	# Ensure site dir exists
	mkdir -p site
	# Generate docs with mkdocs
	mkdocs build
	# Generate API docs
	./hack/api-docs/generate.sh site/genspec.html
	# Add them to spec page originally generated by mkdocs
	sed -i -e '/REPLACE_WITH_GENERATED_CONTENT/{r site/genspec.html' -e 'd}' site/references/spec/index.html
