package src

import (
	"fmt"
	"strconv"
	"strings"
)

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

func transform(word, replace, repalceWith string, occ int) string {
	return strings.Replace(word, replace, repalceWith, occ)
}
