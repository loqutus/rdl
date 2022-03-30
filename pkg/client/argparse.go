package client

import (
	"flag"
	"os"

	"github.com/loqutus/rdl/pkg/types"
)

func ParseArgs() {
	if len(os.Args) < 4 {
		panic("Usage: rdl <actiontype> <filename> <rclonepath> [args]")
	}
	actionType := os.Args[1]
	switch actionType {
	case "torrent":
		types.RunConfig.Mode = "torrent"
	case "youtube":
		types.RunConfig.Mode = "youtube"
	default:
		panic("Unknown action type")
	}
	types.RunConfig.FileName = os.Args[2]
	types.RunConfig.RClonePath = os.Args[3]
	var tempDataDir = flag.String("t", "/tmp/rdl", "temp directory")
	flag.Parse()
	types.RunConfig.TempDataDir = *tempDataDir + "/data"
	types.RunConfig.TempDB = *tempDataDir + "/session.db"
}
