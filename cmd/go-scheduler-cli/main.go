package main

import (
	"context"
	"fmt"
	"github.com/kmacmcfarlane/go-scheduler/internal/grpc"
	"github.com/kmacmcfarlane/go-scheduler/pkg/cli"
	"github.com/kmacmcfarlane/go-scheduler/pkg/common"
	"os"
)

func main() {

	ctx := context.Background()

	logger := common.NewConsoleLogger()

	clientFactory := grpc.NewMasterClientFactory()

	clientService := cli.NewClientService(ctx, clientFactory, logger)

	commandParser := cli.NewCommandParser(clientService, logger)

	// Parse and execute Command
	err := commandParser.Parse(os.Args)

	if nil != err {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}