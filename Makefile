
.PHONY: test

all: | deps build gen_mocks test

deps:
	@echo syncing dependencies...
	@go get github.com/onsi/ginkgo
	@go get github.com/onsi/gomega
	@go get github.com/vektra/mockery/.../
	@go get github.com/stretchr/testify/mock
	@go get google.golang.org/grpc

build: build_cli

build_cli:
	@echo building cli...
	@go build -o go-scheduler-cli ./cmd/go-scheduler-cli

gen_mocks:
	@echo generating mocks...
	@mockery -all -dir pkg/cli -output test/cli/mocks -case=underscore
	@mockery -all -dir pkg/common -output test/common/mocks -case=underscore

test:
	@echo running tests...
	ginkgo -r ./test