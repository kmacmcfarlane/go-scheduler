package test

import "io"

// MockReadCloser wraps an io.Reader adding an impotent Close method to make it a ReadCloser for testing
type MockReadCloser struct {
	reader io.Reader
}

func NewMockReadCloser(reader io.Reader) MockReadCloser {
	return MockReadCloser{
		reader: reader}
}

func (rc MockReadCloser) Read(p []byte) (n int, err error) {
	return rc.reader.Read(p)
}

func (rc MockReadCloser) Close() error {
	// do nothing
	return nil
}