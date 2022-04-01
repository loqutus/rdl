package rclone

import (
	"log"
	"os"
	"os/exec"

	"github.com/loqutus/rdl/pkg/types"
)

func Sync() {
	cmd := exec.Command("rclone", "sync", "-P", types.RunConfig.TempDataDir, types.RunConfig.RClonePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
