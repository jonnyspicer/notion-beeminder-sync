package beeminder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client represents a Beeminder API client
type Client struct {
	BaseURL    string
	AuthToken  string
	HTTPClient *http.Client
}

// NewClient creates a new Beeminder API client
func NewClient(baseURL, authToken string) *Client {
	return &Client{
		BaseURL:   baseURL,
		AuthToken: authToken,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// GetUser retrieves information about a user
func (c *Client) GetUser(params GetUserParams) (*User, error) {
	baseURL := fmt.Sprintf("%s/users/%s.json", c.BaseURL, params.Username)

	query := url.Values{}
	query.Set("auth_token", c.AuthToken)

	if params.Associations {
		query.Set("associations", "true")
	} else {
		query.Set("associations", "true") // always true while User.Goals is a []Goal
	}

	if params.DiffSince != nil {
		query.Set("diff_since", strconv.FormatInt(*params.DiffSince, 10))
	}

	if params.Skinny {
		query.Set("skinny", "true")
	}

	if params.DatapointsCount != nil {
		query.Set("datapoints_count", strconv.FormatInt(*params.DatapointsCount, 10))
	}

	url := baseURL + "?" + query.Encode()

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
