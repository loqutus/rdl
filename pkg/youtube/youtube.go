package youtube

import (
	"log"
	"os"
	"os/exec"

	"github.com/loqutus/rdl/pkg/rclone"
	"github.com/loqutus/rdl/pkg/types"
)

func Download(fileName string) {
	err := rclone.Copy(types.RunConfig.RCloneDrive+":/youtube.archive", types.RunConfig.YoutubeArchive)
	if err != nil {
		log.Println("unable to copy youtube archive:", err)
		log.Println("continuing...")
	}
	defer rclone.Copy(types.RunConfig.YoutubeArchive, types.RunConfig.RCloneDrive+":/youtube.archive")
	err = os.Chdir(types.RunConfig.TempDataDir)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("yt-dlp", "-N", "4", "--download-archive", types.RunConfig.YoutubeArchive, fileName, "--exec", "rclone move {} "+types.RunConfig.RClonePath+";rm "+types.RunConfig.TempDataDir+"/*")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal("cmd.Run() failed with", err)
	}
}
