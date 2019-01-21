
.PHONY: test

all: | deps build gen_mocks test

deps:
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
	@go get github.com/vektra/mockery/.../

build: build_cli

build_cli:
	go build -o go-scheduler-cli ./cmd/go-scheduler-cli

gen_mocks:
	@mockery -all -dir pkg/cli -output test/cli/mocks -case=underscore

test:
	echo "Running tests..."
	ginkgo -r ./test