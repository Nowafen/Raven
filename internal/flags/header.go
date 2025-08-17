package flags

import (
    "strings"
)

var HeaderSlice []string

// ParseHeaders parses header flag values
func ParseHeaders(headers []string) map[string]string {
    result := make(map[string]string)
    for _, h := range headers {
        parts := strings.SplitN(h, ":", 2)
        if len(parts) == 2 {
            result[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
        }
    }
    return result
}
