package rclone

import (
	"os"
	"os/exec"
)

func Copy(sourcePath, destinationPath string) error {
	cmd := exec.Command("rclone", "copy", sourcePath, destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
