package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-builder/internal/disk"
	"github.com/hardenedbsd/hardenedbsd-builder/internal/xz"
)

var (
	release string
)

func main() {
	if err := curl.Run(release); err != nil {
		abort("error: %v\n", err)
	}
	if err := xz.Run(); err != nil {
		abort("error: %v\n", err)
	}
	if err := disk.Mount(); err != nil {
		abort("error: %v\n", err)
	}
	if err := disk.InstallFiles(); err != nil {
		abort("error: %v\n", err)
	}
	if err := disk.Unmount(); err != nil {
		abort("error: %v\n", err)
	}
	if err := xz.Compress(); err != nil {
		abort("error: %v\n", err)
	}
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}

func init() {
	flag.StringVar(&release, "r", "", "The release. Options: 16-CURRENT, 15-STABLE")
	flag.Parse()
}
