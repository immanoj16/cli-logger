package alert

import (
	"fmt"

	"github.com/gookit/color"
)

// Success used to show success alert
func Success(message string, details ...string) string {
	msg := color.Success.Sprint("√ ", message, "\n")
	for _, detail := range details {
		msg += color.Primary.Sprint("- " + detail + "\n")
	}

	return msg
}

// Error used to show error alerts
func Error(message string, err ...error) string {
	msg := color.Danger.Sprint("X ", message, "\n")
	for _, e := range err {
		msg += color.Secondary.Sprint("  " + e.Error() + "\n")
	}

	return msg
}

// Warning used to show warning alerts
func Warning(message string, details ...string) string {
	msg := color.Warn.Sprint("! ", message, "\n")
	for _, detail := range details {
		msg += fmt.Sprint("  " + detail + "\n")
	}

	return msg
}

// Info used to show info alert
func Info(message string) string {
	msg := color.Primary.Sprint("* ", message, "\n")
	return msg
}
