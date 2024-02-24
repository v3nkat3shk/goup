package src

type APIResponse struct {
	Version string `json:"version"`
	Stable  bool   `json:"stable"`
	Files   []File `json:"files"`
}

type File struct {
	Filename string `json:"filename"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Version  string `json:"version"`
	Sha256   string `json:"sha256"`
	Size     int64  `json:"size"`
	Kind     string `json:"kind"`
}

type LocalInstallation struct {
	Installed bool
	Version   string
	Os        string
	Arch      string
}

type Update interface {
	CheckForUpdates() (bool, error)
	DownloadLatestVersion() error
}

type Versions struct {
	LatestVersion APIResponse
	LocalVersion  LocalInstallation
}
