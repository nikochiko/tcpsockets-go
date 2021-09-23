package util

import (
	"fmt"
	"os"
)

// CheckError logs the error if it isn't nil, and exits with status code 1
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
