package urlshortener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const bitlyAPIURL = "https://api-ssl.bitly.com/v4"

// URLShortener struct to hold the access token
type URLShortener struct {
	AccessToken string
}

// ShortenURL shortens a long URL using the Bitly API
func (us *URLShortener) ShortenURL(longURL string) (string, error) {
	url := bitlyAPIURL + "/shorten"
	payload := map[string]string{"long_url": longURL}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+us.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

 // OMG!	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error: %s", body)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["link"].(string), nil
}

// ExpandURL expands a shortened URL using the Bitly API
