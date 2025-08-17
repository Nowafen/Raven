<div align="center">
  
  <a href="https://github.com/Nowafen/Raven/blob/main/README.md">`English`</a> •
  <a href="https://github.com/Nowafen/Raven/blob/main/README.fa.md">`Persian`</a> 
  
</div>
## Raven - Subdomain Discovery Tool

Raven is a modern, efficient, and lightweight subdomain enumeration tool designed for security researchers and penetration testers. Built in Go, it offers a robust and user-friendly experience for discovering subdomains of a target domain, with advanced features to minimize false positives and optimize performance.

---

### Why Raven?

Raven stands out compared to traditional subdomain enumeration techniques due to its modern design and enhanced functionality.

#### Advantages Over Traditional Techniques

##### Optimized Performance with Concurrent Requests
- Uses a worker pool with configurable threads (`default: 10, max: 300`) to send HTTP requests concurrently, reducing scan times drastically.
- Implements rate-limiting (`golang.org/x/time/rate`) to prevent overwhelming target servers.

##### False Positive Mitigation
- Filters out false positives, such as IIS default responses (`<title>IIS Windows Server</title>`).
- Supports customizable status code filtering (`--filter-status`) and matching (`--match-code`).

##### Flexible Wordlist Management
- Supports custom wordlists via `--wordlist`.
- Automatically downloads a default wordlist if none is provided.
- Validates and cleans wordlists, removing invalid entries (e.g., empty lines, trailing dots).

##### Advanced Configuration
- Supports custom HTTP methods (`--method`), headers (`--header`), and proxies (`--proxy`).
- Validation mode (`--validation`) provides quick analysis via status codes.

##### Self-Updating and Version Checking
- `--update` checks for and installs the latest version directly from GitHub.

---

### Why It's More Efficient
Unlike older sequential DNS-based tools, Raven leverages:
- Concurrent HTTP requests
- Intelligent filtering
- Clean wordlist handling
- Color-coded, structured output  

This makes it **faster, cleaner, and more reliable** for real-world penetration testing.

---

### Logic and Workflow

#### 1. Flag Parsing
- Flags parsed with `spf13/pflag`.
- `--domain` is required.
- Other flags (`--wordlist`, `--threads`, `--method`, etc.) have sensible defaults.

#### 2. Wordlist Handling
- If no wordlist is provided, downloads default from:  
  `https://raw.githubusercontent.com/Nowafen/Raven/refs/heads/main/wordlist.txt`
- Cleans invalid entries automatically.

#### 3. Subdomain Scanning
- Generates candidates by combining wordlist entries with the target domain.
- Uses concurrent HTTP client requests with rate-limiting.
- Filters out false positives and irrelevant status codes.

#### 4. Result Output
- Default: prints discovered subdomains.
- With validation: includes color-coded status codes.
- Results can be saved via `--output`.

---

### Installation

#### From Source
```bash
git clone https://github.com/Nowafen/Raven.git
cd Raven
go build -o raven ./cmd/raven
```

#### From Go
```bash
go install github.com/Nowafen/Raven/cmd/raven@latest
```

---

### Usage

Run Raven with the required `--domain` flag and optional parameters:

```bash
raven --domain example.com
```

#### Available Flags

| Flag             | Shorthand | Description                                   | Default                      |
|------------------|-----------|-----------------------------------------------|------------------------------|
| `--domain`       | `-d`      | Target domain to scan (required)              | —                            |
| `--wordlist`     | `-w`      | Path to wordlist file                         | `/tmp/.raven/wordlist.txt`    |
| `--header`       | `-H`      | Custom HTTP headers                           | —                            |
| `--method`       | `-m`      | HTTP method                                   | `GET`                        |
| `--output`       | `-o`      | Output file to save results                   | —                            |
| `--silent`       | —         | Silent mode (no banner/progress)              | `false`                      |
| `--filter-status`| `-f`      | Filter status codes (comma-separated)         | —                            |
| `--match-code`   | `-c`      | Match specific status codes (comma-separated) | —                            |
| `--proxy`        | —         | Proxy URL                                     | —                            |
| `--threads`      | `-t`      | Number of concurrent threads (max: 300)       | `10`                         |
| `--validation`   | `-v`      | Show status codes in output                   | `false`                      |
| `--update`       | —         | Update to latest version                      | —                            |
| `--version`      | —         | Show tool version                             | —                            |
| `--help`         | —         | Show help message                             | —                            |

---

### Example Commands

**Basic Scan**
```bash
raven -d example.com
```

**Scan with Validation**
```bash
raven -d example.com -v
```

**Custom Wordlist & Threads**
```bash
raven -d example.com -w wordlist.txt -t 50 -v
```

**Filter Specific Status Codes**
```bash
raven -d example.com -f 404,403 
```

**Save Results to File**
```bash
raven -d example.com -o results.txt
```

**Silent Mode**
```bash
raven -d example.com --silent
```

**Update Tool**
```bash
raven --update
```

---

### Contributing
Contributions are welcome! [GitHub repository](https://github.com/Nowafen/Raven).

---

## License
Raven is licensed under the **BSD 3-Clause License**. See the [LICENSE](LICENSE) file for details.
