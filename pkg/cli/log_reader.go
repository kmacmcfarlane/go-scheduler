package cli

import (
	"bytes"
	"errors"
	"github.com/kmacmcfarlane/go-scheduler/gen/protobuf/master"
	"io"
)

var _ io.ReadCloser = &LogReader{}

// LogReader exposes the grpc log stream as an io.Reader interface
type LogReader struct {
	logClient master.Master_LogClient
	closer io.Closer
	buffer    *bytes.Buffer
}
 // NewLogReader wraps the logClient and closer implementing ReadCloser and unboxing log messages from the grpc messages
func NewLogReader(logClient master.Master_LogClient, closer io.Closer) *LogReader {
	return &LogReader{
		logClient: logClient,
		closer: closer,
		buffer: new(bytes.Buffer)}
}

// Read handles filling the buffer from the LogClient
func (lr *LogReader) Read(p []byte) (n int, err error){

	if lr.buffer.Len() == 0 {

		resp, err := lr.logClient.Recv()

		if err != nil {
			return n, err
		}

		if resp.Error != "" {
			return n, errors.New(resp.Error)
		}

		lr.buffer.WriteString(resp.LogMessages)
	}

	return lr.buffer.Read(p)
}

// Close closes the underlying connection
func (lr *LogReader) Close() error {
	return lr.closer.Close()
}