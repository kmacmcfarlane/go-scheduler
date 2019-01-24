
.PHONY: test

all: | clean deps gen_protobuf build gen_mocks test

clean:
	@rm -rf ./gen/*
	@rm -rf ./test/cli/mocks/*
	@rm -rf ./test/common/mocks/*

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
	@chmod +x go-scheduler-cli

gen_mocks:
	@echo generating mocks...
	@${GOPATH}/bin/mockery -all -dir pkg/cli -output test/cli/mocks -case=underscore
	@${GOPATH}/bin/mockery -all -dir pkg/common -output test/common/mocks -case=underscore
	@${GOPATH}/bin/mockery -all -dir pkg/master -output test/master/mocks -case=underscore
	@${GOPATH}/bin/mockery -all -dir gen/protobuf/master -output test/master/mocks -case=underscore

gen_protobuf:
	@echo generating protobuf APIs...
	@command -v protoc >/dev/null 2>&1 || { echo >&2 "protoc is not installed. "; exit 1; }
	@protoc --go_out=plugins=grpc,paths=source_relative:./gen ./protobuf/master/master.proto

test:
	@echo running tests...
	@${GOPATH}/bin/ginkgo -r ./test