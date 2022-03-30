package types

type Config struct {
	Mode        string
	TempDB      string
	TempDataDir string
	RClonePath  string
	FileName    string
}

var RunConfig Config
