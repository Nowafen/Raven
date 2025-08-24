package http

import (
    "crypto/tls"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/Nowafen/Raven/internal/flags"
)

// NewClient creates a new HTTP client with the given configuration
func NewClient(cfg flags.Config) (*http.Client, error) {
    transport := &http.Transport{
        TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
        MaxIdleConns:        1000, // Increased for high concurrency
        MaxIdleConnsPerHost: 200,  // Increased for more connections per host
        IdleConnTimeout:     90 * time.Second,
        MaxConnsPerHost:     500,  // New: limit max connections per host
    }
    if cfg.Proxy != "" {
        proxyURL, err := http.ProxyFromEnvironment(&http.Request{URL: nil})
        if err != nil {
            return nil, err
        }
        transport.Proxy = http.ProxyURL(proxyURL)
    }

    client := &http.Client{
        Transport: transport,
        Timeout:   10 * time.Second, // Reduced timeout for faster failures
    }
    return client, nil
}

// MakeRequest sends an HTTP request to the subdomain and returns status code, status, body, and error
func MakeRequest(client *http.Client, cfg flags.Config, subdomain string) (int, string, []byte, error) {
    url := fmt.Sprintf("https://%s", subdomain)
    req, err := http.NewRequest(cfg.Method, url, nil)
    if err != nil {
        return 0, "", nil, err
    }

    for key, value := range cfg.Headers {
        req.Header.Set(key, value)
    }

    resp, err := client.Do(req)
    if err != nil {
        return 0, "", nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return resp.StatusCode, resp.Status, nil, err
    }

    return resp.StatusCode, resp.Status, body, nil
}
