package utils

import "fmt"

// LogError - Responsible for logging error of the application
func LogError(method string, err error) {
	fmt.Print("An error has occurred while"+method+"event", err)
}
