package common

import (
	"bytes"
	"fmt"
	"os"
)

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Write(p []byte) (n int, err error)
}

type ConsoleLogger struct {}

func NewLogger() *ConsoleLogger {
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

	return fmt.Fprint(os.Stderr, buf.String())
}