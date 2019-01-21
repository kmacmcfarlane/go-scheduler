package go_scheduler_cli

import (
	"github.com/kmacmcfarlane/go-scheduler/pkg/cli"
	"os"
)

func main() {

	clientService := cli.NewClientService()

	commandParser := cli.NewCommandParser(clientService)

	// Separate logic for testability
	commandParser.Parse(os.Args)
}