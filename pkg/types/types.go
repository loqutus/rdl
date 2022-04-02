package types

type Config struct {
	Mode           string
	TempDB         string
	TempDataDir    string
	RClonePath     string
	RCloneDrive    string
	FileName       string
	YoutubeArchive string
}

var RunConfig Config
