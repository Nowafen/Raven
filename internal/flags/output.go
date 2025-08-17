package flags

import (
    "fmt"
    "os"
)

// Result holds the scan result for a subdomain
type Result struct {
    Subdomain  string
    StatusCode int
    Status     string
}

// OutputResults outputs the scan results to stdout and optionally to a file
func OutputResults(results []Result, cfg Config) {
    // ANSI color codes
    const (
        ColorReset  = "\033[0m"
        ColorGreen  = "\033[92m" // 200-299
        ColorYellow = "\033[93m" // 300-399
        ColorRed    = "\033[91m" // 400-499
        ColorPurple = "\033[95m" // 500-599
        ColorCyan   = "\033[96m" // 100-199
    )

    var outputFile *os.File
    if cfg.Output != "" {
        var err error
        outputFile, err = os.Create(cfg.Output)
        if err != nil {
            fmt.Fprintf(os.Stderr, "[*] Error creating output file: %v\n", err)
            return
        }
        defer outputFile.Close()
    }

    for _, result := range results {
        var output string
        if cfg.Validation {
            var coloredStatus string
            switch {
            case result.StatusCode >= 100 && result.StatusCode < 200:
                coloredStatus = fmt.Sprintf("%s[%d]%s", ColorCyan, result.StatusCode, ColorReset)
            case result.StatusCode >= 200 && result.StatusCode < 300:
                coloredStatus = fmt.Sprintf("%s[%d]%s", ColorGreen, result.StatusCode, ColorReset)
            case result.StatusCode >= 300 && result.StatusCode < 400:
                coloredStatus = fmt.Sprintf("%s[%d]%s", ColorYellow, result.StatusCode, ColorReset)
            case result.StatusCode >= 400 && result.StatusCode < 500:
                coloredStatus = fmt.Sprintf("%s[%d]%s", ColorRed, result.StatusCode, ColorReset)
            case result.StatusCode >= 500 && result.StatusCode < 600:
                coloredStatus = fmt.Sprintf("%s[%d]%s", ColorPurple, result.StatusCode, ColorReset)
            default:
                coloredStatus = fmt.Sprintf("[%d]", result.StatusCode)
            }
            output = fmt.Sprintf("%s %s\n", result.Subdomain, coloredStatus)
        } else {
            output = fmt.Sprintf("%s\n", result.Subdomain)
        }

        if !cfg.Silent {
            fmt.Print(output)
        }
        if outputFile != nil {
            if cfg.Validation {
                outputFile.WriteString(fmt.Sprintf("%s [%d]\n", result.Subdomain, result.StatusCode))
            } else {
                outputFile.WriteString(fmt.Sprintf("%s\n", result.Subdomain))
            }
        }
    }
}
