package grpc

import (
	gen_master "github.com/kmacmcfarlane/go-scheduler/gen/protobuf/master"
	"github.com/kmacmcfarlane/go-scheduler/pkg/master"
	"google.golang.org/grpc"
	"io"
)

var _ master.MasterClientFactory = MasterClientFactory{}

type MasterClientFactory struct {}

func NewMasterClientFactory() MasterClientFactory{
	return MasterClientFactory{}
}

func (cf MasterClientFactory) CreateClient(host string) (client gen_master.MasterClient, closer io.Closer, err error) {

	//TODO: dial options
	opts := []grpc.DialOption{}

	conn, err := grpc.Dial(host, opts...)

	if err != nil {
		return client, conn, err
	}

	client = gen_master.NewMasterClient(conn)

	return client, conn, err
}