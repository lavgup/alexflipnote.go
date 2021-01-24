package alexflipnote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://api.alexflipnote.dev/"

var httpClient = &http.Client{}

// Client contains fields for interacting with the API
type Client struct {
	Token string
}

// NewClient creates a new client for use with the API
func NewClient(token string) *Client {
	if token == "" {
		panic("Token required.")
	}

	client := &Client{
		Token: token,
	}

	return client
}

// MakeRequest makes a request to the API
func (client *Client) MakeRequest(endpoint string, params url.Values) (interface{}, error) {
	endpointURL := baseURL + endpoint

	if len(params) > 0 {
		endpointURL = endpointURL + "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", client.Token)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request failed, status code %s", strconv.Itoa(res.StatusCode))
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	content := http.DetectContentType(body)
	// JSON response
	if strings.HasPrefix(content, "text/plain") {
		var data map[string]interface{}

		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		return data, nil
	}

	return body, nil
}
