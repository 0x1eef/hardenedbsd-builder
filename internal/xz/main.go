package xz

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

func Decompress(archive string) (string, error) {
	dest := strings.TrimSuffix(archive, ".xz")
	if _, err := os.Stat(dest); errors.Is(err, os.ErrNotExist) {
		args := []string{"-d", archive}
		return dest, cmd.Run(exec.Command("xz", args...))
	}
	return dest, nil
}

func Compress(img string) error {
	args := []string{"-T", "0", "-z", img}
	return cmd.Run(exec.Command("xz", args...))
}
