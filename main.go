package main

import (
	"log"

	"github.com/v3nkat3shk/goup/src"
)

const (
	LATEST_VERSION_URL string = "https://go.dev/dl/?mode=json"
	DOWNLOAD_URL       string = "https://storage.googleapis.com/golang/"
)

func main() {
	versions, err := src.GetVersions(LATEST_VERSION_URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	updateAble, err := versions.CheckForUpdates()

	if err != nil {
		log.Fatal(err.Error())
	}

	if updateAble != false {
		versions.DownloadLatestVersion()
	}
}
