package cli

import (
	"github.com/kmacmcfarlane/go-scheduler/pkg/model/status"
	"io"
)

type ClientService interface {
	Start(dockerImage string, name string, host string) error
	Stop(name string, host string) error
	Query(name string, host string) (status status.Status, err error)
	Log(name string, host string)  (logReader io.Reader, err error)
}

type GrpcClientService struct {

}

func NewClientService() *GrpcClientService {
	return &GrpcClientService{}
}

func (cs *GrpcClientService) Start(dockerImage string, name string, host string) (err error) {
	return err
}

func (cs *GrpcClientService) Stop(name string, host string) (err error) {
	return err
}

func (cs *GrpcClientService) Query(name string, host string) (status status.Status, err error) {
	return status, err
}

func (cs *GrpcClientService) Log(name string, host string)  (logReader io.Reader, err error) {
	return logReader, err
}