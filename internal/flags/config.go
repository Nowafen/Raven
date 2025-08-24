package flags

import (
	"errors"
	"github.com/spf13/pflag"
)

// Config holds all configuration options
type Config struct {
	Domain       string
	Wordlist     string
	Headers      map[string]string
	Method       string
	Output       string
	Silent       bool
	FilterStatus []int
	MatchCode    []int
	Proxy        string
	Threads      int
	RateLimit    float64
	Validation   bool
	Help         bool
	ShortHelp    bool // Flag for -h
	Update       bool
	Version      bool
}

// ParseFlags parses command-line flags and returns a Config
func ParseFlags() (Config, error) {
	cfg := Config{
		Headers: make(map[string]string),
	}

	// Disable default pflag help handling
	pflag.CommandLine.Init("raven", pflag.ExitOnError)
	pflag.CommandLine.SetOutput(nil)

	pflag.StringVarP(&cfg.Domain, "domain", "d", "", "Target domain to scan (e.g., google.com)")
	pflag.StringVarP(&cfg.Wordlist, "wordlist", "w", "", "Path to wordlist file (optional, defaults to .raven/wordlist.txt)")
	pflag.StringSliceVarP(&HeaderSlice, "header", "H", nil, "Custom HTTP headers (e.g., -H Cookie:abcd)")
	pflag.StringVarP(&cfg.Method, "method", "m", "GET", "HTTP method to use (default: GET)")
	pflag.StringVarP(&cfg.Output, "output", "o", "", "Output file to save results")
	pflag.BoolVar(&cfg.Silent, "silent", false, "Run in silent mode (no banner or progress)")
	pflag.IntSliceVarP(&cfg.FilterStatus, "filter-status", "f", nil, "Filter status codes (e.g., 404,400)")
	pflag.IntSliceVarP(&cfg.MatchCode, "match-code", "c", nil, "Match status codes (e.g., 200,301)")
	pflag.StringVar(&cfg.Proxy, "proxy", "", "Proxy URL (e.g., http://127.0.0.1:8080)")
	pflag.IntVarP(&cfg.Threads, "threads", "t", 10, "Number of concurrent threads (default: 10, max: 1000)")
	pflag.Float64VarP(&cfg.RateLimit, "rate-limit", "r", 100.0, "Max requests per second")
	pflag.BoolVarP(&cfg.Validation, "validation", "v", false, "Show status codes for found subdomains")
	pflag.BoolVarP(&cfg.ShortHelp, "help-short", "h", false, "Show concise help message")
	pflag.BoolVar(&cfg.Help, "help", false, "Show detailed help message")
	pflag.BoolVar(&cfg.Update, "update", false, "Update the binary to the latest version")
	pflag.BoolVar(&cfg.Version, "version", false, "Show the tool version")

	pflag.Parse()

	// Skip validations for help, update, or version flags
	if cfg.ShortHelp || cfg.Help || cfg.Update || cfg.Version {
		return cfg, nil
	}

	// Validate flags for scanning
	if cfg.Domain == "" {
		return cfg, errors.New("for usage and help, use the --help flag")
	}
	if err := ValidateDomain(cfg.Domain); err != nil {
		return cfg, err
	}
	if err := ValidateWordlist(cfg.Wordlist); err != nil {
		return cfg, err
	}
	cfg.Headers = ParseHeaders(HeaderSlice)
	if err := ValidateMethod(cfg.Method); err != nil {
		return cfg, err
	}
	if err := ValidateThreads(cfg.Threads); err != nil {
		return cfg, err
	}
	if cfg.Proxy != "" {
		if err := ValidateProxy(cfg.Proxy); err != nil {
			return cfg, err
		}
	}

	return cfg, nil
}
