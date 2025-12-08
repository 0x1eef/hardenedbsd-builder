package xz

import (
	"errors"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

func Run() error {
	if _, err := os.Stat("image.raw"); errors.Is(err, os.ErrNotExist) {
		args := []string{"-d", "-k", "image.raw.xz"}
		return cmd.Run(exec.Command("xz", args...))
	}
	return nil
}