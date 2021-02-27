package logger

import (
	"fmt"
	"os"

	"github.com/immanoj16/cli-logger/pkg/alert"
)

// Success prints a success message and exit status 0.
func Success(message string, details ...string) {
	msg := alert.Success(message, details...)
	fmt.Print(msg)
	os.Exit(0)
}

// Error prints a success message and exit status 1.
func Error(message string, err ...error) {
	msg := alert.Error(message, err...)
	fmt.Print(msg)
	os.Exit(1)
}

// Warning prints a warning message and exit status 1.
func Warning(message string, details ...string) {
	msg := alert.Warning(message, details...)
	fmt.Print(msg)
}

// Info prints an information message without exit.
func Info(message string) {
	msg := alert.Info(message)
	fmt.Print(msg)
}
