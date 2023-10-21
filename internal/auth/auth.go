package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers
// Example below
// Authorization: ApiKey <token>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if len(val) == 0 {
		return "", errors.New("Missing authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Invalid authorization header")
	}

	return vals[1], nil
}
