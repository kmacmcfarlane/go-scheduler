package cli

import (
	"errors"
	"github.com/kmacmcfarlane/go-scheduler/pkg/cli"
	"github.com/kmacmcfarlane/go-scheduler/pkg/model/status"
	"github.com/kmacmcfarlane/go-scheduler/test"
	"github.com/kmacmcfarlane/go-scheduler/test/cli/mocks"
	common_mocks "github.com/kmacmcfarlane/go-scheduler/test/common/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"strings"
)

var _ = Describe("Command Parser", func(){

	var (
		parser *cli.CommandParser
		clientService *mocks.ClientService
		logger *common_mocks.Logger
	)

	BeforeEach(func(){

		clientService = new(mocks.ClientService)
		logger = new(common_mocks.Logger)

		parser = cli.NewCommandParser(clientService, logger)
	})

	AfterEach(func(){
		clientService.AssertExpectations(GinkgoT())
		logger.AssertExpectations(GinkgoT())
	})

	Describe("Subcommand", func(){
		Context("Missing subcommand", func(){
			It("Prints message and returns error", func(){

				logger.
					On("Println", "start").
					Once().
					On("Println", "stop").
					Once().
					On("Println", "query").
					Once().
					On("Println", "log").
					Once()

				logger.
					On("Write", mock.AnythingOfType("[]uint8")).
					Return(123, nil) // called for usage details

				err := parser.Parse([]string{"foo.exe"})

				Ω(err.Error()).Should(Equal("sub-command is required: start, stop, query, or log"))
			})
		})

		Context("Invalid subcommand", func(){
			It("Prints message and returns error", func(){

				err := parser.Parse([]string{"foo.exe", "jump"})

				Ω(err.Error()).Should(Equal("unrecognized command: jump\nsub-command is required: start, stop, query, or log"))
			})
		})
	})

	Describe("Start Command", func(){

		Context("Missing image name", func(){
			It("Prints usage and returns error", func(){

				logger.
					On("Write", mock.AnythingOfType("[]uint8")).
					Return(123, nil) // called for usage details

				err := parser.Parse([]string{"foo.exe", "start", "-name", "jobName"})

				Ω(err.Error()).Should(Equal("image name is required"))
			})
		})

		Context("Missing job name", func(){
			It("Prints usage and returns error", func(){

				logger.
					On("Write", mock.AnythingOfType("[]uint8")).
					Return(123, nil) // called for usage details

				err := parser.Parse([]string{"foo.exe", "start", "-image", "imageName"})

				Ω(err.Error()).Should(Equal("job name is required"))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return", func(){

				clientService.
					On("Start", "imageName", "jobName", "hostName").
					Return(errors.New("error message")).
					Times(1)

				err := parser.Parse([]string{"foo.exe", "start", "-image", "imageName", "-name", "jobName", "-host", "hostName"})

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Success", func(){
			It("Returns no error", func(){

				clientService.
					On("Start", "imageName", "jobName", "hostName").
					Return(nil).
					Times(1)

				logger.
					On("Println", "job started").
					Times(1)

				err := parser.Parse([]string{"foo.exe", "start", "-image", "imageName", "-name", "jobName", "-host", "hostName"})

				Ω(err).Should(BeNil())
			})
		})
	})

	Describe("Stop Command", func(){

		Context("Missing job name", func(){
			It("Prints usage and returns error", func(){

				logger.
					On("Write", mock.AnythingOfType("[]uint8")).
					Return(123, nil) // called for usage details

				err := parser.Parse([]string{"foo.exe", "stop"})

				Ω(err.Error()).Should(Equal("job name is required"))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return", func(){

				clientService.
					On("Stop", "jobName", "hostName").
					Return(errors.New("error message")).
					Times(1)

				err := parser.Parse([]string{"foo.exe", "stop", "-name", "jobName", "-host", "hostName"})

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Success", func(){
			It("Returns no error", func(){

				clientService.
					On("Stop", "jobName", "hostName").
					Return(nil).
					Times(1)

				logger.
					On("Println", "job stopped").
					Times(1)

				err := parser.Parse([]string{"foo.exe", "stop", "-name", "jobName", "-host", "hostName"})

				Ω(err).Should(BeNil())
			})
		})
	})

	Describe("Query Command", func(){

		Context("Missing job name", func(){
			It("Prints usage and returns error", func(){

				logger.
					On("Write", mock.AnythingOfType("[]uint8")).
					Return(123, nil) // called for usage details

				err := parser.Parse([]string{"foo.exe", "query"})

				Ω(err.Error()).Should(Equal("job name is required"))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Print error and return", func(){

				clientService.
					On("Query", "jobName", "hostName").
					Return(status.Status(""), errors.New("error message")).
					Times(1)

				err := parser.Parse([]string{"foo.exe", "query", "-name", "jobName", "-host", "hostName"})

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Success", func(){
			It("Returns no error", func(){

				clientService.
					On("Query", "jobName", "hostName").
					Return(status.Running, nil).
					Times(1)

				logger.
					On("Printf", "job status: %s\n", status.Running.String()).
					Times(1)

				err := parser.Parse([]string{"foo.exe", "query", "-name", "jobName", "-host", "hostName"})

				Ω(err).Should(BeNil())
			})
		})
	})

	Describe("Log Stream Command", func(){

		Context("Missing job name", func(){
			It("Prints usage and returns error", func(){

				logger.
					On("Write", mock.AnythingOfType("[]uint8")).
					Return(123, nil) // called for usage details

				err := parser.Parse([]string{"foo.exe", "log"})

				Ω(err.Error()).Should(Equal("job name is required"))
			})
		})

		Context("Error returned from ClientService", func(){
			It("Returns error", func(){

				clientService.
					On("Log", "jobName", "hostName").
					Return(nil, errors.New("error message")).
					Times(1)

				err := parser.Parse([]string{"foo.exe", "log", "-name", "jobName", "-host", "hostName"})

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Success", func(){
			It("Returns no error", func(){

				logReader := strings.NewReader("one\ntwo\n")
				readCloser := test.NewMockReadCloser(logReader)

				clientService.
					On("Log", "jobName", "hostName").
					Return(readCloser, nil).
					Times(1)

				logger.
					On("Print", "one\n").
					Times(1)

				logger.
					On("Print", "two\n").
					Times(1)

				err := parser.Parse([]string{"foo.exe", "log", "-name", "jobName", "-host", "hostName"})

				Ω(err).Should(BeNil())
			})
		})
	})
})