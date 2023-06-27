package auth

import (
	"errors"
	"net/http"
	"strings"
)

//extracts API key from headers of HTTP request

// format:
// Authorization: ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No auth info found")
	}

	//split on space ApiKey SPACE {apikey}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}

	//no ApiKey part of header
	if vals[0] != "ApiKey" {
		return "", errors.New("Header not complete")
	}

	//return the API key
	return vals[1], nil

}
