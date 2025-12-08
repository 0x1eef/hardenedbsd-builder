package disk

import (
	"os"
	"os/exec"
	"github.com/hardenedbsd/hardenedbsd-builder/internal/cmd"
)

func Run() error {
	if err := setup();  err != nil {
		return err
	}
	if err := mount();  err != nil {
		return err
	}
	return nil
}

func Install() error {
	commands := [][]string{
		{"mkdir", "-p", "/mnt/home/runner/.ssh"},
		{"cp", "keys/hardenedbsd-runner.pub", "/mnt/home/runner/.ssh/authorized_keys"},
		{"cp", "etc/rc.conf", "/mnt/etc/rc.conf"},
		{"cp", "etc/ssh/sshd_config", "/mnt/etc/ssh/sshd_config"},
		{"cp", "etc/rc.local", "/mnt/etc/rc.local"},
	}
	for _, command := range commands {
		err := cmd.Run(exec.Command(command[0], command[1:]...))
		if err != nil {
			return err
		}
	}
	return nil
}

func Teardown() error {
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

func setup() error {
	args := []string{"-a", "-t", "vnode", "-f", "image.raw", "-u", "0"}
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



