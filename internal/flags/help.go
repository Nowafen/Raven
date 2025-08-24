package flags

import "fmt"

// ShowShortHelp displays a concise help message for -h
func ShowShortHelp() {
	fmt.Print(`
Raven - Subdomain Discovery Tool

Usage: raven [options]

Options:
  -d, --domain {domain}        Target domain to scan (e.g., google.com) [required]
  -w, --wordlist {path}        Path to wordlist file (optional, defaults to /tmp/.raven/wordlist.txt)
  -H, --header {header:value}  Custom HTTP headers (e.g., -H Cookie:abcd)
  -m, --method {method}        HTTP method to use (default: GET)
  -o, --output {path}          Output file to save results
  --silent                     Run in silent mode (no banner or progress)
  -f, --filter-status {codes}  Filter status codes (e.g., 404,400)
  -c, --match-code {codes}     Match status codes (e.g., 200,301)
  --proxy {url}                Proxy URL (e.g., http://127.0.0.1:8080)
  -t, --threads {number}       Number of concurrent threads (default: 10, max: 1000)
  -r, --rate-limit {number}    Max requests per second (default: 100)
  -v, --validation             Show status codes for found subdomains
  --update                     Update the binary to the latest version
  --version                    Show the raven version
  -h                           Show this help message
  --help                       Show detailed help message
`)
}

// ShowFullHelp displays a detailed help message for --help
func ShowFullHelp() {
	fmt.Print(`
Raven - Subdomain Discovery Tool

Raven is a fast and efficient tool for discovering subdomains of a target domain using a wordlist.
It supports concurrent scanning, custom HTTP headers, proxy settings, and response filtering.

Usage: raven [options]

Options:
  -d, --domain <domain>        Target domain to scan (e.g., google.com) [required]
                               Specifies the domain to enumerate subdomains for.
                               Example: raven -d example.com

  -w, --wordlist <path>        Path to wordlist file (optional, defaults to /tmp/.raven/wordlist.txt)
                               Wordlist containing subdomain prefixes. Must be a .txt file.
                               Example: raven -d example.com -w /path/to/wordlist.txt
  -H, --header <header:value>  Custom HTTP headers (e.g., -H Cookie:abcd)
                               Add custom headers to HTTP requests. Can be used multiple times.
                               Example: raven -d example.com -H "Cookie: session=abcd" -H "User-Agent: Raven"
  -m, --method <method>        HTTP method to use (default: GET)
                               Supported methods: GET, POST, HEAD, etc.
                               Example: raven -d example.com -m HEAD
  -o, --output <path>          Output file to save results
                               Save found subdomains to a file.
                               Example: raven -d example.com -o results.txt
  --silent                     Run in silent mode (no banner or progress)
                               Suppresses all output except results.
                               Example: raven -d example.com --silent
  -f, --filter-status <codes>  Filter status codes (e.g., 404,400)
                               Exclude subdomains with specified HTTP status codes.
                               Example: raven -d example.com -f 404,403
  -c, --match-code <codes>     Match status codes (e.g., 200,301)
                               Only include subdomains with specified HTTP status codes.
                               Example: raven -d example.com -c 200,301
  --proxy <url>                Proxy URL (e.g., http://127.0.0.1:8080)
                               Route HTTP requests through a proxy.
                               Example: raven -d example.com --proxy http://127.0.0.1:8080
  -t, --threads <number>       Number of concurrent threads (default: 10, max: 1000)
                               Controls the number of concurrent HTTP requests.
                               Example: raven -d example.com -t 100
  -r, --rate-limit <number>    Max requests per second (default: 100)
                               Limits the rate of HTTP requests to avoid overwhelming the server.
                               Example: raven -d example.com -r 500
  -v, --validation             Show status codes for found subdomains
                               Displays HTTP status codes in the output for valid subdomains.
                               Example: raven -d example.com -v
  --update                     Update the binary to the latest version
                               Checks and updates Raven to the latest release.
                               Example: raven --update
  --version                    Show the tool version
                               Displays the current version of Raven.
                               Example: raven --version
  -h                           Show concise help message
                               Displays a brief overview of available options.
  --help                       Show this detailed help message

Examples:
  Basic scan with default settings:
    raven -d example.com
  Scan with custom wordlist and high concurrency:
    raven -d example.com -w wordlist.txt -t 200 -r 500
  Scan with proxy and filtered status codes:
    raven -d example.com --proxy http://127.0.0.1:8080 -f 404,403
  Silent mode with output to file:
    raven -d example.com --silent -o results.txt
  Scan with custom headers and validation:
    raven -d example.com -H "Cookie: session=abcd" -v

Notes:
  - Ensure the wordlist is a .txt file with one subdomain prefix per line.
  - High thread counts or rate limits may trigger server rate-limiting or IP blocking; use proxies if needed.
  - Use --silent for automation to reduce output noise.
  - Run --update periodically to stay up-to-date with the latest features.
`)
}
