package curl

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

var (
	urls = map[string]string{
		"16-CURRENT": "https://github.com/0x1eef/hardenedbsd-builder/releases/download/16CURRENT_UFS_AMD64_ORIGINAL/hardenedbsd-vm.raw.xz",
		"15-STABLE":  "https://github.com/0x1eef/hardenedbsd-builder/releases/download/15STABLE_UFS_AMD64_ORIGINAL/hardenedbsd-vm.raw.xz",
	}
)

func Run(release string) error {
	var (
		url string
		ok  bool
	)
	if url, ok = urls[release]; !ok {
		return fmt.Errorf("unknown release (%s)", release)
	}
	if _, err := os.Stat("image.raw.xz"); errors.Is(err, os.ErrNotExist) {
		args := []string{"-L", "-o", "image.raw.xz", url}
		return cmd.Run(exec.Command("curl", args...))
	}
	return nil
}
