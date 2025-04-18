package logs

import "log"

type ErrorLogger struct {
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{}
}

func (e *ErrorLogger) ConsoleLogError(err error) {
	// Log the error (for example, print it to the console)
	// In a real application, you might want to log this to a file or monitoring system
	log.Println("Error:", err.Error())
}
