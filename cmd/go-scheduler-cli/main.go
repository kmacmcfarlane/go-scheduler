package go_scheduler_cli

import (
	"github.com/kmacmcfarlane/go-scheduler/pkg/cli"
	"github.com/kmacmcfarlane/go-scheduler/pkg/common"
	"os"
)

func main() {

	clientService := cli.NewClientService()

	logger := common.NewLogger()

	commandParser := cli.NewCommandParser(clientService, logger)

	// Parse and Execute Command
	statusCode := commandParser.Parse(os.Args)

	os.Exit(statusCode)
}