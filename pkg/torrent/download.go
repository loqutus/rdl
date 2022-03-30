package torrent

import (
	"log"
	"os"
	"time"

	rainTorrent "github.com/cenkalti/rain/torrent"
	"github.com/loqutus/rdl/pkg/types"
)

func Download() {
	torrentFile, err := os.Open(types.RunConfig.FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer torrentFile.Close()
	config := rainTorrent.DefaultConfig
	config.DataDir = types.RunConfig.TempDataDir
	config.Database = types.RunConfig.TempDB
	config.DataDirIncludesTorrentID = false
	ses, err := rainTorrent.NewSession(config)
	if err != nil {
		log.Fatal(err)
	}
	defer ses.Close()
	torrentOptions := rainTorrent.AddTorrentOptions{
		StopAfterDownload: true,
	}
	tor, err := ses.AddTorrent(torrentFile, &torrentOptions)
	if err != nil {
		log.Fatal(err)
	}
	for range time.Tick(time.Second) {
		s := tor.Stats()
		if s.Status.String() == "Stopped" {
			break
		}
		log.Printf("Status: %s, Downloaded: %d, Peers: %d", s.Status.String(), s.Bytes.Completed, s.Peers.Total)
	}
}
