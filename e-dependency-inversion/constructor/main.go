package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type MyService struct {
	logger Logger // Dependency
}

func NewMyService(logger Logger) *MyService { // Constructor injection
	return &MyService{logger: logger}
}

func (s *MyService) DoSomething() {
	s.logger.Log("MyService is doing something")
}

func main() {
	logger := ConsoleLogger{}
	service := NewMyService(logger) // Injecting the logger
	service.DoSomething()
}
