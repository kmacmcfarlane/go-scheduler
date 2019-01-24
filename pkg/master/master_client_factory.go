package master

import (
	"github.com/kmacmcfarlane/go-scheduler/gen/protobuf/master"
	"io"
)

type MasterClientFactory interface {
	CreateClient(host string) (result master.MasterClient, closer io.Closer, err error)
}