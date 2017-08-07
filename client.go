package awesomeminer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client is a consumer of the Awesome Miner HTTP API.
// Username and Password are used for HTTP Basic Auth (if available).
type Client struct {
	URL      string
	Username string
	Password string
}

// NewClient creates a new instance of Client with which API calls to
// Awesome Miner can be made.
func NewClient(url, username, password string) Client {
	return Client{URL: url, Username: username, Password: password}
}

func (c *Client) newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	u := fmt.Sprintf("%s/api/%s", c.URL, endpoint)
	req, err := http.NewRequest(method, u, body)
	if c.Username != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}

	return req, err
}

func (c *Client) doGetRequest(endpoint string, v interface{}) (interface{}, error) {
	req, err := c.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyText, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
