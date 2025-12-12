package curl

import (
	"errors"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

const (
	image = "https://github.com/0x1eef/hardenedbsd-builder/releases/download/16CURRENT_UFS_AMD64_ORIGINAL/hardenedbsd-vm.raw.xz"
)

func Run() error {
	if _, err := os.Stat("image.raw.xz"); errors.Is(err, os.ErrNotExist) {
		args := []string{"-L", "-o", "image.raw.xz", url()}
		return cmd.Run(exec.Command("curl", args...))
	}
	return nil
}

func url() string {
	return image
}
