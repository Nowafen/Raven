package flags

import (
    "errors"
    "fmt"
    "os"
)

// ValidateThreads validates the threads flag
func ValidateThreads(threads int) error {
    if threads < 1 {
        return errors.New("threads must be at least 1")
    }
    if threads > 1000 {
        return errors.New("maximum allowed threads is 1000")
    }
    if threads > 500 {
        fmt.Printf("[*] Critical Warning: Using %d threads may significantly impact CPU and network performance. Proceed? [y/N] ", threads)
        var input string
        fmt.Scanln(&input)
        if input != "y" && input != "Y" {
            os.Exit(1)
        }
    } else if threads > 200 {
        fmt.Printf("[*] Warning: Using %d threads may cause high CPU/network usage. Proceed? [y/N] ", threads)
        var input string
        fmt.Scanln(&input)
        if input != "y" && input != "Y" {
            os.Exit(1)
        }
    }
    return nil
}
