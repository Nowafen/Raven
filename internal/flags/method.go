package flags

import (
    "errors"
    "strings"
)

// ValidateMethod validates the HTTP method
func ValidateMethod(method string) error {
    validMethods := []string{"GET", "POST", "HEAD", "TRACE"}
    for _, m := range validMethods {
        if strings.ToUpper(method) == m {
            return nil
        }
    }
    return errors.New("invalid HTTP method")
}
