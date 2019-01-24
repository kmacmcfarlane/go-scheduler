package cli

import (
	"context"
	"errors"
	gen_master "github.com/kmacmcfarlane/go-scheduler/gen/protobuf/master"
	"github.com/kmacmcfarlane/go-scheduler/pkg/common"
	"github.com/kmacmcfarlane/go-scheduler/pkg/master"
	"github.com/kmacmcfarlane/go-scheduler/pkg/model/status"
	"io"
)

type ClientService interface {
	Start(dockerImage string, jobName string, host string) error
	Stop(jobName string, host string) error
	Query(jobName string, host string) (status status.Status, err error)
	Log(jobName string, host string)  (logReader io.ReadCloser, err error)
}

type GrpcClientService struct {
	ctx context.Context
	clientFactory master.MasterClientFactory
	logger common.Logger
}

func NewClientService(
	ctx context.Context,
	clientFactory master.MasterClientFactory,
	logger common.Logger) *GrpcClientService {

	return &GrpcClientService{
		ctx: ctx,
		clientFactory: clientFactory,
		logger: logger}
}

// Start creates or starts a job and assigns a name to it
func (cs *GrpcClientService) Start(dockerImage string, name string, host string) (err error) {

	client, conn, err := cs.clientFactory.CreateClient(host)

	if err != nil {
		return err
	}

	defer conn.Close()

	req := &gen_master.StartRequest{
		DockerImage: dockerImage,
		JobName: name}

	resp, err := client.Start(cs.ctx, req)

	if err != nil {
		return err
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return err
}

// Stop terminates a job
func (cs *GrpcClientService) Stop(name string, host string) (err error) {

	client, conn, err := cs.clientFactory.CreateClient(host)

	if err != nil {
		return err
	}

	defer conn.Close()

	req := &gen_master.StopRequest{
		JobName: name}

	resp, err := client.Stop(cs.ctx, req)

	if err != nil {
		return err
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return err
}

// Query provides the current status of the given job
func (cs *GrpcClientService) Query(name string, host string) (result status.Status, err error) {

	client, conn, err := cs.clientFactory.CreateClient(host)

	if err != nil {
		return result, err
	}

	defer conn.Close()

	req := &gen_master.QueryRequest{
		JobName: name}

	resp, err := client.Query(cs.ctx, req)

	if err != nil {
		return result, err
	}

	if resp.Error != "" {
		return result, errors.New(resp.Error)
	}

	result = status.New(resp.Status)

	return result, err
}

// Log provides a reader that streams the log output of the job
func (cs *GrpcClientService) Log(name string, host string)  (logReader io.ReadCloser, err error) {

	client, conn, err := cs.clientFactory.CreateClient(host)

	if err != nil {
		return logReader, err
	}

	req := &gen_master.LogRequest{
		JobName: name}

	logClient, err := client.Log(cs.ctx, req)

	if err != nil {
		return logReader, err
	}

	logReader = NewLogReader(logClient, conn)

	return logReader, err
}
