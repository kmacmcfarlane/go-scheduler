
.PHONY: test

all: | clean deps build gen_mocks gen_protobuf test

clean:
	@rm -rf ./pkg/gen/*

deps:
	@echo syncing dependencies...
	@go get github.com/onsi/ginkgo
	@go get github.com/onsi/gomega
	@go get github.com/vektra/mockery/.../
	@go get github.com/stretchr/testify/mock
	@go get -u google.golang.org/grpc
	@go get -u github.com/golang/protobuf/protoc-gen-go

build: build_cli

build_cli:
	@echo building cli...
	@go build -o go-scheduler-cli ./cmd/go-scheduler-cli

gen_mocks:
	@echo generating mocks...
	@mockery -all -dir pkg/cli -output test/cli/mocks -case=underscore
	@mockery -all -dir pkg/common -output test/common/mocks -case=underscore

gen_protobuf:
	@echo generating protobuf APIs...
	@command -v protoc >/dev/null 2>&1 || { echo >&2 "protoc is not installed. "; exit 1; }
	@protoc --go_out=paths=source_relative:./pkg/gen ./protobuf/master/.proto

test:
	@echo running tests...
	@ginkgo -r ./test