package common

import (
	"bytes"
	"fmt"
)

// Wrap common output functions for testability
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Write(p []byte) (n int, err error)
}

var _ Logger = &ConsoleLogger{}

// ConsoleLogger implements Logger interface and logs to Stdout
type ConsoleLogger struct {}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Print(args ...interface{}) {
	fmt.Print(args)
}

func (l *ConsoleLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (l *ConsoleLogger) Println(args ...interface{}) {
	fmt.Println(args)
}

// io.Writer interface
func (l *ConsoleLogger) Write(p []byte) (n int, err error) {

	buf := bytes.NewBuffer(p)

	return fmt.Print(buf.String())
}