package disk

import (
	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
	"os"
	"os/exec"
)

func Mount(img string) error {
	if err := setup(img); err != nil {
		return err
	}
	if err := mount(); err != nil {
		return err
	}
	return nil
}

func InstallFiles() error {
	commands := [][]string{
		{"mkdir", "-p", "/mnt/root/.ssh"},
		{"cp", "config/keys/id_ed25519.pub", "/mnt/root/.ssh/authorized_keys"},
		{"cp", "config/etc/rc.conf", "/mnt/etc/rc.conf"},
		{"cp", "config/etc/ssh/sshd_config", "/mnt/etc/ssh/sshd_config"},
		{"cp", "config/etc/rc.local", "/mnt/etc/rc.local"},
		{"cp", "config/boot/loader.conf", "/mnt/boot/loader.conf"},
	}
	for _, command := range commands {
		err := cmd.Run(exec.Command(command[0], command[1:]...))
		if err != nil {
			return err
		}
	}
	return nil
}

func Unmount() error {
	commands := [][]string{
		{"umount", "/mnt"},
		{"mdconfig", "-d", "-u", "0"},
	}
	for _, command := range commands {
		err := cmd.Run(exec.Command(command[0], command[1:]...))
		if err != nil {
			return err
		}
	}
	return nil
}

func setup(img string) error {
	args := []string{"-a", "-t", "vnode", "-f", img, "-u", "0"}
	if err := cmd.Run(exec.Command("mdconfig", args...)); err != nil {
		return err
	}
	return nil
}

func mount() error {
	if err := os.MkdirAll("/mnt", 0o755); err != nil {
		return err
	} else {
		args := []string{"/dev/md0p4", "/mnt"}
		return cmd.Run(exec.Command("mount", args...))
	}
}
