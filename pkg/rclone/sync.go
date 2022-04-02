package rclone

import (
	"log"
	"os"
	"os/exec"
)

func Sync(sourcePath, destinationPath string) {
	cmd := exec.Command("rclone", "sync", "-P", sourcePath, destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
