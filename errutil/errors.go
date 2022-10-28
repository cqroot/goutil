package errutil

import (
	"fmt"
	"os"
)

var logFunc func(error) = func(err error) {
	fmt.Fprint(os.Stderr, fmt.Sprintf("ERROR: %s\n", err))
}

func SetLogFunc(f func(error)) {
	logFunc = f
}

func LogWithError(err error) {
	logFunc(err)
}

func LogIfError(err error) {
	if err != nil {
		LogWithError(err)
	}
}

func ExitWithError(err error) {
	LogWithError(err)
	os.Exit(1)
}

func ExitIfError(err error) {
	if err != nil {
		ExitWithError(err)
	}
}
