package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

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

func (v Versions) CheckForUpdates() (bool, error) {
	return v.canBeUpdated()
}

func (v Versions) DownloadLatestVersion() error {
	return nil
}

func GetVersions(url string) (Update, error) {
	latestVerison, err := getLatestVersion(url)

	if err != nil {
		return nil, err
	}

	localVersion, err := getLocalVersion()

	if err != nil {
		return nil, err
	}

	return Versions{
		LatestVersion: latestVerison,
		LocalVersion:  localVersion,
	}, nil
}

func getLatestVersion(url string) (APIResponse, error) {
	data := make([]APIResponse, 1)
	res, err := http.Get(url)
	if err != nil {
		return APIResponse{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return APIResponse{}, err

	}

	if err := json.Unmarshal(body, &data); err != nil {
		return APIResponse{}, err
	}

	return data[0], nil
}

func getLocalVersion() (LocalInstallation, error) {
	goVersion, err := exec.Command("go", "version").Output()

	fmt.Println("loacal go version: ", string(goVersion))

	if err != nil {
		return LocalInstallation{
			Installed: false,
			Version:   "",
			Os:        "",
			Arch:      "",
		}, err
	}

	return LocalInstallation{
		Installed: true,
		Version:   "go1.21.7",
		Os:        "linux",
		Arch:      "amd64",
	}, nil
}

func (v *Versions) canBeUpdated() (bool, error) {
	latestVersion, err := convertVerion(strings.Split(strings.Replace(v.LatestVersion.Version, "go", "", 1), "."))
	if err != nil {
		return false, err
	}

	localVersion, err := convertVerion(strings.Split(strings.Replace(v.LocalVersion.Version, "go", "", 1), "."))
	if err != nil {
		return false, err
	}

	return shouldUpdate(localVersion, latestVersion)
}

func convertVerion(versionArray []string) ([]uint, error) {

	userVersionInt := make([]uint, len(versionArray))

	for idx, value := range versionArray {
		intValue, err := strconv.ParseUint(value, 0, 32)
		if err != nil {
			return nil, err
		}
		userVersionInt[idx] = uint(intValue)
	}

	return userVersionInt, nil

}

func shouldUpdate(local, latest []uint) (bool, error) {
	if len(local) != len(latest) {
		return false, fmt.Errorf("cannot be compair local and latest go version")
	}

	for idx, value := range latest {
		if local[idx] >= value {
			return false, nil
		}
		if value > local[idx] {
			return true, nil
		}
	}

	return false, nil

}
