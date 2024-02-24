package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (v Versions) CheckForUpdates() (bool, error) {
	return v.canBeUpdated()
}

func (v Versions) DownloadLatestVersion() error {
	version := v.LatestVersion.Version
	os := v.LocalVersion.Os
	arch := v.LocalVersion.Arch

	finalVersion := fmt.Sprintf("%s %s/%s", version, os, arch)
	fmt.Println("Version to Download: ", finalVersion)
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
