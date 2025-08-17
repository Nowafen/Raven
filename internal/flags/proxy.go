package flags

import "errors"

// ValidateProxy validates the proxy URL
func ValidateProxy(proxy string) error {
    if proxy == "" {
        return errors.New("proxy URL cannot be empty")
    }
    return nil
}
