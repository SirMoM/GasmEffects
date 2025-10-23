package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Location returns an absolute path of the caller file and the line number.
func Location() string {
	_, file, line, _ := runtime.Caller(3)
	p, _ := os.Getwd()
	location := fmt.Sprintf("%s:%d", strings.TrimPrefix(file, p), line)
	return strings.Replace(location, "/Users/i13az81/dev/lern/go/go-wasm/", "", 1)
}

var verbose bool = true
var writer = os.Stdout

func log(obj any, verbose bool, level string) {
	var lvl, loc, msg string

	if level != "" {
		lvl = level
	}
	if verbose {
		loc = Location() + "\n"
	}
	if _, ok := obj.(fmt.Stringer); ok {
		msg = fmt.Sprintf("%s%s %s\n", loc, lvl, obj)
	} else {
		msg = fmt.Sprintf("%s%s %+v\n", loc, lvl, obj)
	}
	_, _ = fmt.Fprintln(writer, msg)
}

func info(obj any) {
	log(obj, verbose, "[INFO]")
}

func warn(obj any) {
	log(obj, verbose, "[WARN]")
}

func ERR(obj any) {
	writer = os.Stderr
	log(obj, verbose, "[ERROR]")
	writer = os.Stdout
}
