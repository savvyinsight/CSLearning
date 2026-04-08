package main

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Log(level string, message string)
}

type FileLogger struct {
	file *os.File
}

func (f *FileLogger) Log(level, message string) {
	fmt.Fprintf(f.file, "[%s] %s: %s\n",
		time.Now().Format(time.RFC3339),
		level,
		message)
}

type ConsoleLogger struct {
	prefix string
}

func (c *ConsoleLogger) Log(level, message string) {
	fmt.Printf("%s [%s] %s\n",
		c.prefix,
		level,
		message)
}

// service using the logger
type UserService struct {
	logger Logger
}

func (s UserService) CreateUser(name string) error {
	s.logger.Log("INFO", fmt.Sprintf("Creating user: %s", name))
	return nil
}

func createLogFile() *os.File {
	file, _ := os.Create("app.log")
	return file
}

func main() {
	fileLogger := &FileLogger{file: createLogFile()}
	userService := UserService{logger: fileLogger}
	userService.CreateUser("Alice")

	consoleLogger := &ConsoleLogger{prefix: "USER-SERVICE"}
	userService.logger = consoleLogger
	userService.CreateUser("Bob")
}
