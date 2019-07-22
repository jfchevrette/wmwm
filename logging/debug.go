package logging

import (
	"fmt"
	"os"
)

var Debug bool

func Println(args ...interface{}) {
	if !Debug {
		return
	}
	fmt.Fprintln(os.Stderr, args...)
}

func Print(args ...interface{}) {
	if !Debug {
		return
	}
	fmt.Fprint(os.Stderr, args...)
}

func Fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func Error(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
