package flags

import "fmt"

// ShowHelp displays the help message
func ShowHelp() {
    fmt.Print(`
Raven - Subdomain Discovery Tool

Usage: raven [options]

Options:
  -d, --domain <domain>        Target domain to scan (e.g., google.com) [required]
  -w, --wordlist <path>        Path to wordlist file (optional, defaults to .raven/wordlist.txt)
  -H, --header <header:value>  Custom HTTP headers (e.g., -H Cookie:abcd)
  -m, --method <method>        HTTP method to use (default: GET)
  -o, --output <path>          Output file to save results
  --silent                     Run in silent mode (no banner or progress)
  -f, --filter-status <codes>  Filter status codes (e.g., 404,400)
  -c, --match-code <codes>     Match status codes (e.g., 200,301)
  --proxy <url>                Proxy URL (e.g., http://127.0.0.1:8080)
  -t, --threads <number>       Number of concurrent threads (default: 10, max: 300)
  -v, --validation             Show status codes for found subdomains
  --update                     Update the binary to the latest version
  --version                    Show the tool version
  --help                       Show this help message
`)
