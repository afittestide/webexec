package main

import (
	"bufio"
	"fmt"
	"github.com/pion/webrtc/v3"
	"os"
	"regexp"
)

// TokensFilePath holds the path to a file where each authorized token has a line
var TokensFilePath = ConfPath("authorized_tokens")

// ReadAuthorizedTokens reads the tokens file and returns all the tokens in it
func ReadAuthorizedTokens() ([]string, error) {
	var tokens []string
	file, err := os.Open(TokensFilePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open authorized_tokens: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed to read authorized_tokens: %s", err)
	}
	return tokens, nil
}

// IsAuthorized checks whether a client token is authorized
func IsAuthorized(token string) bool {
	tokens, err := ReadAuthorizedTokens()
	if err != nil {
		Logger.Error(err)
		return false
	}
	for _, at := range tokens {
		if token == at {
			return true
		}
	}
	return false
}

// GetFingerprint extract the fingerprints from a client's offer
func GetFingerprint(offer *webrtc.SessionDescription) string {
	r, _ := regexp.Compile("a=fingerprint:.+ ([0-9A-Z]{2}:)+[0-9A-Z]{2}")
	fp := r.FindString(offer.SDP)
	Logger.Infof("fingerprint=%s sdp=%s", fp, offer.SDP)
	if len(fp) > 14 {
		return fp[14:]
	}
	return ""
}
