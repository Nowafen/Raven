package discovery

import (
    "bufio"
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "sync"
    "github.com/Nowafen/Raven/internal/flags"
    "github.com/Nowafen/Raven/internal/http"
    "golang.org/x/time/rate"
)

// ScanSubdomains scans subdomains using a worker pool
func ScanSubdomains(cfg flags.Config, words []string) ([]flags.Result, error) {
    // Create temporary file for subdomains
    tmpFile, err := ioutil.TempFile("", "raven-subdomains-*.txt")
    if err != nil {
        return nil, err
    }
    defer os.Remove(tmpFile.Name())

    // Write subdomains to temporary file in a streaming manner
    fileWriter, err := os.OpenFile(tmpFile.Name(), os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }
    defer fileWriter.Close()
    writer := bufio.NewWriter(fileWriter)
    for _, word := range words {
        fmt.Fprintf(writer, "%s.%s\n", word, cfg.Domain)
    }
    writer.Flush()

    client, err := http.NewClient(cfg)
    if err != nil {
        return nil, err
    }

    var results []flags.Result
    var mu sync.Mutex
    jobs := make(chan string, 100000)
    resultsChan := make(chan flags.Result, 100000)
    limiter := rate.NewLimiter(rate.Limit(cfg.RateLimit), int(cfg.RateLimit))
    ctx := context.Background()

    var wg sync.WaitGroup
    for i := 0; i < cfg.Threads; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for subdomain := range jobs {
                if err := limiter.Wait(ctx); err != nil {
                    continue
                }
                statusCode, status, body, err := http.MakeRequest(client, cfg, subdomain)
                if err != nil {
                    continue
                }
                if strings.Contains(string(body), "<title>IIS Windows Server</title>") {
                    continue
                }
                if shouldInclude(statusCode, cfg) {
                    mu.Lock()
                    resultsChan <- flags.Result{
                        Subdomain:  fmt.Sprintf("https://%s", subdomain),
                        StatusCode: statusCode,
                        Status:     status,
                    }
                    mu.Unlock()
                }
            }
        }()
    }

    file, err := os.Open(tmpFile.Name())
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
    for scanner.Scan() {
        jobs <- scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    close(jobs)
    wg.Wait()
    close(resultsChan)

    for res := range resultsChan {
        results = append(results, res)
    }
    return results, nil
}

// shouldInclude checks if the status code matches the filter/match criteria
func shouldInclude(statusCode int, cfg flags.Config) bool {
    if statusCode == 0 {
        return false
    }
    if len(cfg.FilterStatus) > 0 {
        for _, code := range cfg.FilterStatus {
            if statusCode == code {
                return false
            }
        }
    }
    if len(cfg.MatchCode) > 0 {
        for _, code := range cfg.MatchCode {
            if statusCode == code {
                return true
            }
        }
        return false
    }
    return true
}
