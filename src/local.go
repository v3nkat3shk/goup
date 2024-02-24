package src

import (
	"os/exec"
	"strings"
)

func getLocalVersion() (LocalInstallation, error) {
	goVersion, err := exec.Command("go", "version").Output()

	if err != nil {
		return LocalInstallation{Installed: false}, err
	}

	version := strings.Split(strings.TrimSpace(transform(string(goVersion), "go version", "", 1)), " ")
	osInfo := strings.Split(version[1], "/")

	if err != nil {
		return LocalInstallation{}, err
	}

	return LocalInstallation{
		Installed: true,
		Version:   version[0],
		Os:        osInfo[0],
		Arch:      osInfo[1],
	}, nil
}
