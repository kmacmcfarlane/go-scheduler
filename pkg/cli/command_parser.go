package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/kmacmcfarlane/go-scheduler/pkg/common"
	"io"
)

type CommandParser struct {
	clientService ClientService
	logger common.Logger
}

func NewCommandParser(clientService ClientService, logger common.Logger) *CommandParser {
	return &CommandParser{
		clientService: clientService,
		logger: logger}
}

func (cp *CommandParser) Parse(args []string) (err error) {

	// Start
	startCommand := flag.NewFlagSet("start", flag.ContinueOnError)
	startCommand.SetOutput(cp.logger)

	startImage := startCommand.String("image", "", "The docker image name (Required)")
	startName := startCommand.String("name", "", "The name of the job (Required)")
	startHost := startCommand.String("host", "localhost", "The hostname of the master node")
	// Improvement: port number flag

	// Stop
	stopCommand := flag.NewFlagSet("stop", flag.ContinueOnError)
	stopCommand.SetOutput(cp.logger)

	stopName := stopCommand.String("name", "", "The name of the job (Required)")
	stopHost := stopCommand.String("host", "localhost", "The hostname of the master node")

	// Query
	queryCommand := flag.NewFlagSet("query", flag.ContinueOnError)
	queryCommand.SetOutput(cp.logger)

	queryName := queryCommand.String("name", "", "The name of the job (Required)")
	queryHost := queryCommand.String("host", "localhost", "The hostname of the master node")

	// Stream Logs
	logCommand := flag.NewFlagSet("log", flag.ContinueOnError)
	logCommand.SetOutput(cp.logger)

	logName := logCommand.String("name", "", "The name of the job (Required)")
	logHost := logCommand.String("host", "localhost", "The hostname of the master node")

	// Validate command input
	defaultFlags := flag.NewFlagSet("default", flag.ContinueOnError)
	helpShort := defaultFlags.Bool("h", false, "Print this usage")
	helpLong := defaultFlags.Bool("help", false, "Print this usage")

	defaultFlags.Parse(args)

	if len(args) < 2 || *helpShort || *helpLong {

		cp.logger.Println("start")
		startCommand.PrintDefaults()
		cp.logger.Println("stop")
		stopCommand.PrintDefaults()
		cp.logger.Println("query")
		queryCommand.PrintDefaults()
		cp.logger.Println("log")
		logCommand.PrintDefaults()

		return errors.New("sub-command is required: start, stop, query, or log")
	}

	// Parse sub-command
	switch args[1] {
	case "start":

		err = startCommand.Parse(args[2:])

		if err != nil {
			return err
		}

		// Assert required flags
		if *startImage == "" {

			startCommand.PrintDefaults()
			return errors.New("image name is required")
		}

		if  *startName == "" {

			startCommand.PrintDefaults()
			return errors.New("job name is required")
		}

		// Call master node
		err := cp.clientService.Start(*startImage, *startName, *startHost)

		if nil != err {
			return err
		} else {
			cp.logger.Println("job started")
		}
	case "stop":

		stopCommand.Parse(args[2:])

		// Assert required flags
		if *stopName == "" {
			stopCommand.PrintDefaults()
			return errors.New("job name is required")
		}

		// Call master node
		err := cp.clientService.Stop(*stopName, *stopHost)

		if nil != err {
			return err
		} else {
			cp.logger.Println("job stopped")
		}
	case "query":

		queryCommand.Parse(args[2:])

		// Assert required flags
		if *queryName == "" {
			queryCommand.PrintDefaults()
			return errors.New("job name is required")
		}

		// Call master node
		status, err := cp.clientService.Query(*queryName, *queryHost)

		if nil != err {
			return err
		}

		cp.logger.Printf("job status: %s\n", status.String())
	case "log":

		logCommand.Parse(args[2:])

		// Assert required flags
		if *logName == "" {
			logCommand.PrintDefaults()
			return errors.New("job name is required")
		}

		// Call master node
		logReader, err := cp.clientService.Log(*logName, *logHost)

		if nil != err {
			return err
		}

		defer logReader.Close()

		// Print out the streaming log data
		bufferedReader := bufio.NewReader(logReader)

		for {

			line, err := bufferedReader.ReadString('\n')

			if err != nil {
				if err == io.EOF {
					break
				} else {
					return err
				}
			}

			cp.logger.Print(line) // line contains the newline char at the end
		}
	default:
		return errors.New(fmt.Sprintf("unrecognized command: %s\nsub-command is required: start, stop, query, or log", args[1]))
	}

	return err
}