package main

import (
	"fmt"
	"os"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-builder/internal/disk"
	"github.com/hardenedbsd/hardenedbsd-builder/internal/xz"
)

func main() {
	if err := curl.Run(); err != nil {
		abort("error: %v\n", err)
	}
	if err := xz.Run(); err != nil {
		abort("error: %v\n", err)
	}
	if err := disk.Run(); err != nil {
		abort("error: %v\n", err)
	}
	if err := disk.Install(); err != nil {
		abort("error: %v\n", err)
	}
	if err := disk.Teardown(); err != nil {
		abort("error: %v\n", err)
	}
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}
