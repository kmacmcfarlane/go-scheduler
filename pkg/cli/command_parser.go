package cli

import (
	"bufio"
	"flag"
	"io"
	"os"
)

type CommandParser struct {
	clientService ClientService
	logger Logger
}

func NewCommandParser(clientService ClientService, logger Logger) *CommandParser {
	return &CommandParser{
		clientService: clientService,
		logger: logger}
}

func (cp *CommandParser) Parse(args []string) (statusCode int) {

	// Start
	startCommand := flag.NewFlagSet("start", flag.ExitOnError)

	startImage := startCommand.String("image", "", "The docker image name (Required)")
	startName := startCommand.String("name", "", "The name of the job (Required)")
	startHost := startCommand.String("host", "localhost", "The hostname of the master node")
	// Improvement: port number flag

	// Stop
	stopCommand := flag.NewFlagSet("stop", flag.ExitOnError)

	stopName := stopCommand.String("name", "", "The name of the job (Required)")
	stopHost := stopCommand.String("host", "localhost", "The hostname of the master node")

	// Query
	queryCommand := flag.NewFlagSet("query", flag.ExitOnError)

	queryName := queryCommand.String("name", "", "The name of the job (Required)")
	queryHost := queryCommand.String("host", "localhost", "The hostname of the master node")

	// Stream Logs
	logCommand := flag.NewFlagSet("log", flag.ExitOnError)

	logName := logCommand.String("name", "", "The name of the job (Required)")
	logHost := logCommand.String("host", "localhost", "The hostname of the master node")

	// Validate command input
	if len(args) < 2 {
		cp.logger.Println("sub-command is required: start, stop, query, or log")
		flag.PrintDefaults()
		return 1
	}

	// Parse sub-command
	switch args[1] {
	case "start":
		startCommand.Parse(args[2:])

		// Assert required flags
		if *startImage == "" || *startName == "" {
			flag.PrintDefaults()
			return 1
		}

		// Call master node
		err := cp.clientService.Start(*startImage, *startName, *startHost)

		if nil != err {
			cp.logger.Printf("error starting job: %s", err.Error())
			return 3
		} else {
			cp.logger.Println("job started")
		}
	case "stop":
		stopCommand.Parse(args[2:])

		// Assert required flags
		if *stopName == "" {
			flag.PrintDefaults()
			return 1
		}

		// Call master node
		err := cp.clientService.Stop(*stopName, *stopHost)

		if nil != err {
			cp.logger.Printf("error stopping job: %s", err.Error())
			return 3
		} else {
			cp.logger.Println("job stopped")
		}
	case "query":
		queryCommand.Parse(args[2:])

		// Assert required flags
		if *queryName == "" {
			flag.PrintDefaults()
			return 1
		}

		// Call master node
		status, err := cp.clientService.Query(*queryName, *queryHost)

		if nil != err {
			cp.logger.Printf("error querying job: %s", err.Error())
			return 3
		}

		cp.logger.Printf("job status: %s\n", status.String())
	case "log":
		logCommand.Parse(args[2:])

		// Assert required flags
		if *logName == "" {
			flag.PrintDefaults()
			return 1
		}

		// Call master node
		logReader, err := cp.clientService.Log(*logName, *logHost)

		if nil != err {
			cp.logger.Printf("error streaming logs from job: %s", err.Error())
			return 3
		}

		// Print out the streaming log data
		bufferedReader := bufio.NewReader(logReader)

		for {
			line, err := bufferedReader.ReadString('\n')

			if err != nil {
				if err != io.EOF {
					cp.logger.Printf("error streaming logs from job: %s", err.Error())
					os.Exit(3)
				} else {
					break
				}
			}

			cp.logger.Print(line) // line contains the newline char at the end
		}
	default:
		cp.logger.Printf("unrecognized command: %s\n", args[1])
		return 1
	}

	return 0
}