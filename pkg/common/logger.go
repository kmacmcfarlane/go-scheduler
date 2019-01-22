package common

import "fmt"

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
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