package flags

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Version is the current version of the Raven tool
const Version = "0.1.3"

// VersionURL is the URL to check the latest version and binary URL
const VersionURL = "https://raw.githubusercontent.com/Nowafen/Raven/refs/heads/main/Version.json"

// VersionInfo holds the version and binary URL from version.json
type VersionInfo struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

// CheckVersion checks if the current version is outdated
func CheckVersion() (bool, error) {
	latestVersion, err := getLatestVersion()
	if err != nil {
		return false, err
	}
	return isOutdated(Version, latestVersion), nil
}

// ShowVersion returns the latest version and checks if the current version is outdated
func ShowVersion() (string, error) {
	return getLatestVersion()
}

// GetVersionInfo fetches the version and binary URL from VersionURL
func GetVersionInfo() (VersionInfo, error) {
	var versionInfo VersionInfo

	resp, err := http.Get(VersionURL)
	if err != nil {
		return versionInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return versionInfo, errors.New("failed to fetch version file")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return versionInfo, err
	}

	if err := json.Unmarshal(body, &versionInfo); err != nil {
		return versionInfo, errors.New("failed to parse version.json: " + err.Error())
	}

	if versionInfo.Version == "" || versionInfo.URL == "" {
		return versionInfo, errors.New("invalid version.json format")
	}

	return versionInfo, nil
}

// getLatestVersion fetches the latest version from VersionURL
func getLatestVersion() (string, error) {
	versionInfo, err := GetVersionInfo()
	if err != nil {
		return "", err
	}
	return versionInfo.Version, nil
}

// isOutdated compares two version strings (format: x.y.z)
func isOutdated(current, latest string) bool {
	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	// Ensure both versions have 3 parts (x.y.z)
	for i := 0; i < 3; i++ {
		var currentNum, latestNum int
		if i < len(currentParts) {
			currentNum = parseInt(currentParts[i])
		}
		if i < len(latestParts) {
			currentNum = parseInt(latestParts[i])
		}
		if currentNum < latestNum {
			return true
		}
		if currentNum > latestNum {
			return false
		}
	}
	return false
}

// parseInt converts a string to int, returns 0 if invalid
func parseInt(s string) int {
	var result int
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		result = result*10 + int(r-'0')
	}
	return result
}
