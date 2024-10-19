package urlshortener



import (

	"bytes"

	"encoding/json"

	"fmt"
 // 5
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



	if resp.StatusCode != http.StatusOK {

		body, _ := ioutil.ReadAll(resp.Body)

		return "", fmt.Errorf("error: %s", body)

	}



	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result["link"].(string), nil

}



// ExpandURL expands a shortened URL using the Bitly API

func (us *URLShortener) ExpandURL(shortURL string) (string, error) {

	url := bitlyAPIURL + "/expand"

	payload := map[string]string{"bitlink_id": shortURL}

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



	if resp.StatusCode != http.StatusOK {

		body, _ := ioutil.ReadAll(resp.Body)

		return "", fmt.Errorf("error: %s", body)

	}



	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result["long_url"].(string), nil

}



// GetClicks retrieves the number of clicks for a shortened URL

func (us *URLShortener) GetClicks(shortURL string) (int, error) {

	url := bitlyAPIURL + "/bitlinks/" + shortURL + "/clicks"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return 0, err

	}

	req.Header.Set("Authorization", "Bearer "+us.AccessToken)



	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		return 0, err

	}

	defer resp.Body.Close()



	if resp.StatusCode != http.StatusOK {

		body, _ := ioutil.ReadAll(resp.Body)

		return 0, fmt.Errorf("error: %s", body)

	}



	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return int(result["link_clicks"].(float64)), nil

}



