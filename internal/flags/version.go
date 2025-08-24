package flags

import (
    "errors"
    "io"
    "net/http"
    "strings"
)

// Version is the current version of the Raven tool
const Version = "0.1.0"

// VersionURL is the URL to check the latest version
const VersionURL = "https://github.com/Nowafen/Raven/raw/refs/heads/main/Version"

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

// getLatestVersion fetches the latest version from VersionURL
func getLatestVersion() (string, error) {
    resp, err := http.Get(VersionURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("failed to fetch version file")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return strings.TrimSpace(string(body)), nil
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
            latestNum = parseInt(latestParts[i])
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
