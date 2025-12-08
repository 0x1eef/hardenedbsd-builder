package curl

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

const (
	image = "https://download.freebsd.org/releases/VM-IMAGES/%s/%s/Latest/FreeBSD-%s-%s-%s.raw.xz"
)

func Run() error {
	if _, err := os.Stat("image.raw.xz"); errors.Is(err, os.ErrNotExist) {
		args := []string{"-L", "-o", "image.raw.xz", url()}
		return cmd.Run(exec.Command("curl", args...))
	}
	return nil
}

func url() string {
	return fmt.Sprintf(
		image,
		"15.0-RELEASE", "amd64",
		"15.0-RELEASE", "amd64",
		"ufs",
	)
}
