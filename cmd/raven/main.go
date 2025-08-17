package main

import (
    "fmt"
    "os"
    "time"
    "github.com/Nowafen/Raven/internal/discovery"
    "github.com/Nowafen/Raven/internal/flags"
)

func main() {
    // Parse flags
    cfg, err := flags.ParseFlags()
    if err != nil {
        fmt.Fprintf(os.Stderr, "[\033[96m*\033[0m] Error: %v\n", err)
        os.Exit(1)
    }

    // Check if no flags are provided
    if len(os.Args) == 1 {
        fmt.Println("[\033[96m*\033[0m] No flags provided. Use --help for usage information.")
        os.Exit(1)
    }

    // Handle version flag
    if cfg.Version {
        latestVersion, err := flags.ShowVersion()
        if err != nil {
            fmt.Fprintf(os.Stderr, "[\033[96m*\033[0m] Error checking version: %v\n", err)
        } else if latestVersion == flags.Version {
            fmt.Printf("[\033[96m*\033[0m] Raven is up to date (version: \033[92m%s\033[0m)\n", flags.Version)
        } else {
            fmt.Printf("[\033[96m*\033[0m] \033[92mA new version of Raven is available: \033[92m%s\033[0m (current: \033[91m%s\033[0m)\033[0m\n", latestVersion, flags.Version)
        }
        os.Exit(0)
    }

    // Handle help flag
    if cfg.Help {
        flags.ShowHelp()
        os.Exit(0)
    }

    // Handle update flag
    if cfg.Update {
        if err := flags.UpdateBinary(); err != nil {
            fmt.Fprintf(os.Stderr, "[\033[96m*\033[0m] Error updating binary: %v\n", err)
            os.Exit(1)
        }
        fmt.Println("[\033[96m*\033[0m] Binary updated successfully")
        os.Exit(0)
    }

    // Show banner unless silent mode is enabled
    if !cfg.Silent {
        fmt.Print(`
    ____                       
   / __ \ ____ __  _____  ____ 
  / /_/ / __  / | / / _ \/ __ \
 / _, _/ /_/ /| |/ /  __/ / / /
/_/ |_|\__,_/ |___/\___/_/ /_/ 
                        Version:` + flags.Version + `

`)
    }

    // Check version unless in silent mode
    if !cfg.Silent {
        outdated, err := flags.CheckVersion()
        if err == nil && outdated {
            fmt.Println("[\033[96m*\033[0m] \033[92mA new version of Raven is available! Run 'raven --update' to update.\033[0m")
        }
    }

    // Show configuration unless silent mode is enabled
    if !cfg.Silent {
        fmt.Printf("[\033[96m*\033[0m] Target Domain: %s\n", cfg.Domain)
        if cfg.Wordlist != "" {
            fmt.Printf("[\033[96m*\033[0m] Wordlist: %s\n", cfg.Wordlist)
        } else {
            fmt.Printf("[\033[96m*\033[0m] Wordlist: %s (default)\n", flags.DefaultWordlistPath)
        }
        fmt.Printf("[\033[96m*\033[0m] HTTP Method: %s\n", cfg.Method)
        fmt.Printf("[\033[96m*\033[0m] Threads: %d\n", cfg.Threads)
        if cfg.Proxy != "" {
            fmt.Printf("[\033[96m*\033[0m] Proxy: %s\n", cfg.Proxy)
        }
        if len(cfg.FilterStatus) > 0 {
            fmt.Printf("[\033[96m*\033[0m] Filtering Status Codes: %v\n", cfg.FilterStatus)
        }
        if len(cfg.MatchCode) > 0 {
            fmt.Printf("[\033[96m*\033[0m] Matching Status Codes: %v\n", cfg.MatchCode)
        }
        if len(cfg.Headers) > 0 {
            fmt.Printf("[\033[96m*\033[0m] Headers: %v\n", cfg.Headers)
        }
        fmt.Printf("[\033[96m*\033[0m] Validation: %v\n", cfg.Validation)
    }

    // Validate wordlist and only show messages if cleaned
    _, err = flags.ReadWordlist(cfg.Wordlist, cfg.Silent)
    if err != nil && err.Error() != "wordlist was cleaned due to invalid entries" {
        fmt.Fprintf(os.Stderr, "[\033[96m*\033[0m] Error: %v\n", err)
        os.Exit(1)
    }

    // Start timer
    startTime := time.Now()

    if !cfg.Silent {
        fmt.Println("[\033[96m*\033[0m] Scanning started...")
    }

    // Start subdomain discovery
    results, err := discovery.ScanSubdomains(cfg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "[\033[96m*\033[0m] Error: %v\n", err)
        os.Exit(1)
    }

    // Calculate duration
    duration := time.Since(startTime).Minutes()

    // Output results
    flags.OutputResults(results, cfg)

    if cfg.Silent {
        return
    }

    if cfg.Output != "" {
        fmt.Printf("[\033[96m*\033[0m] Scanning completed. Results saved to: %s\n", cfg.Output)
    }
    fmt.Printf("[\033[96m*\033[0m] Total subdomains found: %d\n", len(results))
    fmt.Printf("[\033[96m*\033[0m] Scan completed in %.2fm\n", duration)
}
