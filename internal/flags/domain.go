package flags

import (
    "errors"
    "strings"
)

// ValidateDomain validates the domain flag
func ValidateDomain(domain string) error {
    if domain == "" {
        return errors.New("domain is required")
    }
    if !strings.Contains(domain, ".") {
        return errors.New("invalid domain format")
    }
    return nil
}
