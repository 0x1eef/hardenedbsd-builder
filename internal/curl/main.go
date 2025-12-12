package curl

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

var (
	dest = "hardenedbsd-vm.raw.xz"
	base = "https://github.com/0x1eef/hardenedbsd-builder/releases/download/"
	urls = map[string]string{
		"16-CURRENT": fmt.Sprintf("%s/16CURRENT_UFS_AMD64_ORIGINAL/hardenedbsd-vm.raw.xz", base),
		"15-STABLE":  fmt.Sprintf("%s/15STABLE_UFS_AMD64_ORIGINAL/hardenedbsd-vm.raw.xz", base),
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
	if _, err := os.Stat(dest); errors.Is(err, os.ErrNotExist) {
		args := []string{"-L", "-o", dest, url}
		return cmd.Run(exec.Command("curl", args...))
	}
	return nil
}
