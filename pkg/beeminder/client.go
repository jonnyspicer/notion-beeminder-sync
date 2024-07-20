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

// User represents a Beeminder user
type User struct {
	Username     string        `json:"username"`
	Timezone     string        `json:"timezone"`
	UpdatedAt    int64         `json:"updated_at"`
	Goals        []interface{} `json:"goals"` // Can be []string or []Goal depending on parameters
	Deadbeat     bool          `json:"deadbeat"`
	UrgencyLoad  float64       `json:"urgency_load"`
	DeletedGoals []struct {
		ID string `json:"id"`
	} `json:"deleted_goals,omitempty"`
}

// Goal represents a Beeminder goal
type Goal struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Add other goal fields as needed
	Datapoints    []Datapoint `json:"datapoints,omitempty"`
	LastDatapoint *Datapoint  `json:"last_datapoint,omitempty"`
}

// Datapoint represents a Beeminder datapoint
type Datapoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
	Comment   string  `json:"comment"`
	ID        string  `json:"id"`
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

// GetUserParams represents the parameters for the GetUser method
type GetUserParams struct {
	Username        string
	Associations    bool
	DiffSince       *int64
	Skinny          bool
	DatapointsCount *int64
}

// GetUser retrieves information about a user
func (c *Client) GetUser(params GetUserParams) (*User, error) {
	baseURL := fmt.Sprintf("%s/users/%s.json", c.BaseURL, params.Username)

	query := url.Values{}
	query.Set("auth_token", c.AuthToken)

	if params.Associations {
		query.Set("associations", "true")
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
