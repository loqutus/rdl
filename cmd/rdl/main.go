package main

import (
	"github.com/loqutus/rdl/pkg/client"
	"github.com/loqutus/rdl/pkg/files"
	"github.com/loqutus/rdl/pkg/rclone"
	"github.com/loqutus/rdl/pkg/torrent"
	"github.com/loqutus/rdl/pkg/types"
	"github.com/loqutus/rdl/pkg/youtube"
)

func main() {
	client.ParseArgs()
	files.EnsureExists(types.RunConfig.TempDataDir)
	switch types.RunConfig.Mode {
	case "torrent":
		torrent.Download()
		rclone.Sync()
	case "youtube":
		youtube.Download()
	}
}
