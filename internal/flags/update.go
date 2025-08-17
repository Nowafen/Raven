package flags

import (
    "errors"
    "os"
    "os/exec"
    "path/filepath"
)

// UpdateBinary updates the Raven binary to the latest version
func UpdateBinary() error {
    // Get the path of the current binary
    currentBinary, err := os.Executable()
    if err != nil {
        return err
    }

    // Run go install to get the latest version
    cmd := exec.Command("go", "install", "github.com/Nowafen/Raven/cmd/raven@latest")
    cmd.Env = append(os.Environ(), "GOBIN="+filepath.Dir(currentBinary))
    if err := cmd.Run(); err != nil {
        return errors.New("failed to update binary: " + err.Error())
    }

    return nil
}
