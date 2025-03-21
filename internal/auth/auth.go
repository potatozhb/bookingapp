package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("AUthorization")
	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid authentication info found")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Invalid authentication info in first part found")
	}
	return vals[1], nil
}
