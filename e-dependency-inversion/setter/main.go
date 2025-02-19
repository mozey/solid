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

func (s *MyService) SetLogger(logger Logger) { // Setter injection
	s.logger = logger
}

func (s *MyService) DoSomething() {
	if s.logger != nil { // Important to check for nil!
		s.logger.Log("MyService is doing something")
	} else {
		fmt.Println("No logger set!")
	}
}

func main() {
	service := MyService{} // Object created without the dependency
	service.DoSomething()  // No logger yet

	logger := ConsoleLogger{}
	service.SetLogger(logger) // Injecting the logger
	service.DoSomething()     // Now it works
}
