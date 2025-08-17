package flags

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

// DefaultWordlistURL is the URL to download the default wordlist
const DefaultWordlistURL = "https://raw.githubusercontent.com/Nowafen/Raven/refs/heads/main/wordlist.txt"
const DefaultWordlistPath = "/tmp/.raven/wordlist.txt"

// ReadWordlist reads the wordlist from the given path or downloads the default
func ReadWordlist(path string, silent bool) ([]string, error) {
    if path == "" {
        path = DefaultWordlistPath
        if err := downloadDefaultWordlist(silent); err != nil {
            return nil, err
        }
    }

    if !strings.HasSuffix(path, ".txt") {
        return nil, errors.New("wordlist must be a .txt file")
    }

    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        if word != "" {
            words = append(words, word)
        }
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    if len(words) == 0 {
        return nil, errors.New("wordlist is empty")
    }

    // Validate and clean wordlist
    cleanedWords, cleaned := cleanWordlist(words)
    if cleaned && !silent {
        fmt.Println("[\033[96m*\033[0m] Validating wordlist...")
        fmt.Println("[\033[96m*\033[0m] We made the wordlist valid")
        fmt.Println("[\033[96m*\033[0m] Now wordlist is valid")
    }
    return cleanedWords, nil
}

// ValidateWordlist validates the wordlist path
func ValidateWordlist(path string) error {
    if path == "" {
        return nil // Will use default wordlist
    }
    if !strings.HasSuffix(path, ".txt") {
        return errors.New("wordlist must be a .txt file")
    }
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return errors.New("wordlist file does not exist")
    }
    return nil
}

// downloadDefaultWordlist downloads the default wordlist if it doesn't exist
func downloadDefaultWordlist(silent bool) error {
    if _, err := os.Stat(DefaultWordlistPath); err == nil {
        return nil // File already exists
    }

    // Show downloading message unless in silent mode
    if !silent {
        fmt.Println("[\033[96m*\033[0m] Downloading default wordlist...")
    }

    resp, err := http.Get(DefaultWordlistURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("failed to download default wordlist")
    }

    if err := os.MkdirAll(filepath.Dir(DefaultWordlistPath), 0755); err != nil {
        return err
    }

    file, err := os.Create(DefaultWordlistPath)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
    return err
}

// cleanWordlist validates and cleans the wordlist
func cleanWordlist(words []string) ([]string, bool) {
    var cleanedWords []string
    invalidChars := regexp.MustCompile(`[/:*?<>|]`)
    cleaned := false

    for _, word := range words {
        word = strings.TrimSpace(word)
        if word == "" {
            cleaned = true
            continue
        }
        // Remove trailing dot
        if strings.HasSuffix(word, ".") {
            word = strings.TrimSuffix(word, ".")
            cleaned = true
        }
        // Check for invalid characters
        if invalidChars.MatchString(word) {
            cleaned = true
            continue
        }
        cleanedWords = append(cleanedWords, word)
    }
    return cleanedWords, cleaned
}
