package cli

import (
	"bytes"
	"context"
	"errors"
	"github.com/kmacmcfarlane/go-scheduler/gen/protobuf/master"
	"github.com/kmacmcfarlane/go-scheduler/pkg/cli"
	"github.com/kmacmcfarlane/go-scheduler/pkg/model/status"
	common_mocks "github.com/kmacmcfarlane/go-scheduler/test/common/mocks"
	"github.com/kmacmcfarlane/go-scheduler/test/master/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
)

var _ = Describe("Client Service", func(){

	var (
		clientService cli.ClientService
		clientFactory *mocks.MasterClientFactory
		client *mocks.MasterClient
		logClient *mocks.Master_LogClient
		closer *common_mocks.Closer
		ctx context.Context
		logger *common_mocks.Logger
	)

	BeforeEach(func(){

		logger = new(common_mocks.Logger)

		closer = new(common_mocks.Closer)

		logClient = new(mocks.Master_LogClient)

		client = new(mocks.MasterClient)

		clientFactory = new(mocks.MasterClientFactory)

		clientService = cli.NewClientService(ctx, clientFactory, logger)
	})

	AfterEach(func(){
		logger.AssertExpectations(GinkgoT())
	})

	Describe("Start Job", func(){

		Context("Client factory error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, errors.New("error message")).
					Once()

				err := clientService.Start("imageName", "jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client call error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.StartRequest{
					JobName:"jobName",
					DockerImage:"imageName"}

				client.
					On("Start", ctx, request).
					Return(nil, errors.New("error message"))

				closer.
					On("Close").
					Return(nil).
					Once()

				err := clientService.Start("imageName", "jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client response contains error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.StartRequest{
					JobName:"jobName",
					DockerImage:"imageName"}

				response := &master.StartResponse{
					Error:"error message"}

				client.
					On("Start", ctx, request).
					Return(response, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				err := clientService.Start("imageName", "jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("No errors", func(){
			It("Returns no errors and has an empty error field in the response", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.StartRequest{
					JobName:"jobName",
					DockerImage:"imageName"}

				response := &master.StartResponse{
					Error:""}

				client.
					On("Start", ctx, request).
					Return(response, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				err := clientService.Start("imageName", "jobName", "hostName")

				Ω(err).Should(BeNil())
			})
		})
	})

	Describe("Stop Job", func(){

		Context("Client factory error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, errors.New("error message")).
					Once()

				err := clientService.Stop("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client call error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.StopRequest{
					JobName:"jobName"}

				client.
					On("Stop", ctx, request).
					Return(nil, errors.New("error message"))

				closer.
					On("Close").
					Return(nil).
					Once()

				err := clientService.Stop("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client response contains error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.StopRequest{
					JobName:"jobName"}

				response := &master.StopResponse{
					Error:"error message"}

				client.
					On("Stop", ctx, request).
					Return(response, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				err := clientService.Stop("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("No errors", func(){
			It("Returns no errors and has an empty error field in the response", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.StopRequest{
					JobName:"jobName"}

				response := &master.StopResponse{
					Error:""}

				client.
					On("Stop", ctx, request).
					Return(response, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				err := clientService.Stop("jobName", "hostName")

				Ω(err).Should(BeNil())
			})
		})
	})

	Describe("Query Job", func(){

		Context("Client factory error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, errors.New("error message")).
					Once()

				_, err := clientService.Query("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client call error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.QueryRequest{
					JobName:"jobName"}

				client.
					On("Query", ctx, request).
					Return(nil, errors.New("error message"))

				closer.
					On("Close").
					Return(nil).
					Once()

				_, err := clientService.Query("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client response contains error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.QueryRequest{
					JobName:"jobName"}

				response := &master.QueryResponse{
					Error:"error message",
					Status: status.Unknown.String()}

				client.
					On("Query", ctx, request).
					Return(response, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				_, err := clientService.Query("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("No errors", func(){
			It("Returns no errors and has an empty error field in the response", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.QueryRequest{
					JobName:"jobName"}

				response := &master.QueryResponse{
					Error:"",
					Status: status.Running.String()}

				client.
					On("Query", ctx, request).
					Return(response, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				result, err := clientService.Query("jobName", "hostName")

				Ω(err).Should(BeNil())
				Ω(result).Should(Equal(status.Running))
			})
		})
	})

	Describe("Log Stream From Job", func(){

		Context("Client factory error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, errors.New("error message")).
					Once()

				_, err := clientService.Log("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("Client call error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.LogRequest{
					JobName:"jobName"}

				client.
					On("Log", ctx, request).
					Return(nil, errors.New("error message"))

				closer.
					On("Close").
					Return(nil).
					Once()

				_, err := clientService.Log("jobName", "hostName")

				Ω(err.Error()).Should(Equal("error message"))

			})
		})

		Context("Client response contains error", func(){
			It("Returns the error", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.LogRequest{
					JobName:"jobName"}

				response := &master.LogResponse{
					Error:"error message",
					LogMessages: "one\ntwo\n"}

				logClient.
					On("Recv").
					Return(response, nil).
					On("Recv").
					Return(nil, io.EOF)

				client.
					On("Log", ctx, request).
					Return(logClient, nil)

				closer.
					On("Close").
					Return(nil).
					Once()

				reader, err := clientService.Log("jobName", "hostName")

				Ω(err).Should(BeNil())

				defer reader.Close()

				// Drain the reader into the buffer
				buf := new(bytes.Buffer)
				_, err = buf.ReadFrom(reader)

				Ω(err.Error()).Should(Equal("error message"))
			})
		})

		Context("No errors", func(){
			It("Returns no errors and reader gives us one\\ntwo\\nthree\\nfour\\nEOF", func(){

				clientFactory.
					On("CreateClient", "hostName").
					Return(client, closer, nil).
					Once()

				request := &master.LogRequest{
					JobName:"jobName"}

				response := &master.LogResponse{
					Error:"",
					LogMessages: "one\ntwo\n"}

				response2 := &master.LogResponse{
					Error:"",
					LogMessages: "three\nfour\n"}

				logClient.
					On("Recv").
					Return(response, nil).
					Once().
					On("Recv").
					Return(response2, nil).
					Once().
					On("Recv").
					Return(nil, io.EOF).
					Once()

				client.
					On("Log", ctx, request).
					Return(logClient, nil).
					Once()

				closer.
					On("Close").
					Return(nil).
					Once()

				reader, err := clientService.Log("jobName", "hostName")

				defer reader.Close()

				Ω(err).Should(BeNil())

				// Drain the reader into the buffer
				buf := new(bytes.Buffer)
				_, err = buf.ReadFrom(reader)

				Ω(err).Should(BeNil())

				Ω(buf.String()).Should(Equal("one\ntwo\nthree\nfour\n"))
			})
		})
	})
})