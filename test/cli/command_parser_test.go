package cli

import (
	"errors"
	"github.com/kmacmcfarlane/go-scheduler/pkg/cli"
	"github.com/kmacmcfarlane/go-scheduler/pkg/model/status"
	"github.com/kmacmcfarlane/go-scheduler/test/cli/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Command Parser", func(){

	var (
		parser *cli.CommandParser
		clientService *mocks.ClientService
		logger *mocks.Logger
	)

	BeforeEach(func(){

		clientService = new(mocks.ClientService)
		logger = new(mocks.Logger)

		parser = cli.NewCommandParser(clientService, logger)
	})

	AfterEach(func(){

		clientService.AssertExpectations(GinkgoT())

		logger.AssertExpectations(GinkgoT())
	})

	Describe("Subcommand", func(){
		Context("Missing subcommand", func(){
			It("Prints message and exits with error status 1", func(){

				logger.
					On("Println", "sub-command is required: start, stop, query, or log").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe"})

				Ω(statusCode).Should(Equal(1))
			})
		})

		Context("Invalid subcommand", func(){
			It("Prints message and exits with error status 1", func(){

				logger.
					On("Printf", "unrecognized command: %s\n", "jump").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "jump"})

				Ω(statusCode).Should(Equal(1))
			})
		})
	})

	Describe("Start Command", func(){

		Context("Missing image name", func(){
			It("Prints usage and exits with error status 1", func(){

				statusCode := parser.Parse([]string{"foo.exe", "start", "-name", "jobName"})

				Ω(statusCode).Should(Equal(1))
			})
		})

		Context("Missing job name", func(){
			It("Prints usage and exits with error status 1", func(){

				statusCode := parser.Parse([]string{"foo.exe", "start", "-image", "imageName"})

				Ω(statusCode).Should(Equal(1))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return status 3", func(){

				clientService.
					On("Start", "imageName", "jobName", "hostName").
					Return(errors.New("error message")).
					Times(1)

				logger.
					On("Printf", "error starting job: %s", "error message").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "start", "-image", "imageName", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(3))
			})
		})

		Context("Success", func(){
			It("return status 0", func(){

				clientService.
					On("Start", "imageName", "jobName", "hostName").
					Return(nil).
					Times(1)

				logger.
					On("Println", "job started").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "start", "-image", "imageName", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(0))
			})
		})
	})

	Describe("Stop Command", func(){

		Context("Missing job name", func(){
			It("Prints usage and exits with error status 1", func(){

				statusCode := parser.Parse([]string{"foo.exe", "stop"})

				Ω(statusCode).Should(Equal(1))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return status 3", func(){

				clientService.
					On("Stop", "jobName", "hostName").
					Return(errors.New("error message")).
					Times(1)

				logger.
					On("Printf", "error stopping job: %s", "error message").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "stop", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(3))
			})
		})

		Context("Success", func(){
			It("return status 0", func(){

				clientService.
					On("Stop", "jobName", "hostName").
					Return(nil).
					Times(1)

				logger.
					On("Println", "job stopped").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "stop", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(0))
			})
		})
	})

	Describe("Query Command", func(){

		Context("Missing job name", func(){
			It("Prints usage and exits with error status 1", func(){

				statusCode := parser.Parse([]string{"foo.exe", "query"})

				Ω(statusCode).Should(Equal(1))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return status 3", func(){

				clientService.
					On("Query", "jobName", "hostName").
					Return(status.Status(""), errors.New("error message")).
					Times(1)

				logger.
					On("Printf", "error querying job: %s", "error message").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "query", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(3))
			})
		})

		Context("Success", func(){
			It("return status 0", func(){

				clientService.
					On("Query", "jobName", "hostName").
					Return(status.Running, nil).
					Times(1)

				logger.
					On("Printf", "job status: %s\n", status.Running.String()).
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "query", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(0))
			})
		})
	})

	Describe("Log Stream Command", func(){

		Context("Missing job name", func(){
			It("Prints usage and exits with error status 1", func(){

				statusCode := parser.Parse([]string{"foo.exe", "log"})

				Ω(statusCode).Should(Equal(1))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return status 3", func(){

				clientService.
					On("Log", "jobName", "hostName").
					Return(nil, errors.New("error message")).
					Times(1)

				logger.
					On("Printf", "error streaming logs from job: %s", "error message").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "log", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(3))
			})
		})

		Context("Success", func(){
			It("return status 0", func(){

				logReader := strings.NewReader("one\ntwo\n")

				clientService.
					On("Log", "jobName", "hostName").
					Return(logReader, nil).
					Times(1)

				logger.
					On("Print", "one\n").
					Times(1)

				logger.
					On("Print", "two\n").
					Times(1)

				statusCode := parser.Parse([]string{"foo.exe", "log", "-name", "jobName", "-host", "hostName"})

				Ω(statusCode).Should(Equal(0))
			})
		})
	})
})