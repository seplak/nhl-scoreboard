package utils

import (
	"net/http"
	"io"

	log "github.com/sirupsen/logrus"
)

func MakeRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Error fetching: " + url)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read the response body from: " + url)
	}
	return body
}
