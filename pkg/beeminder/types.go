package beeminder

// ============================== //
// ======== BEEMINDER API ======= //
// ============================== //

// User represents a Beeminder user
type User struct {
	Username     string  `json:"username"`
	Timezone     string  `json:"timezone"`
	UpdatedAt    int64   `json:"updated_at"`
	Goals        []Goal  `json:"goals"` // if this ever gets extracted, this needs changing as can also be []string if associations = false in the API call
	Deadbeat     bool    `json:"deadbeat"`
	UrgencyLoad  float64 `json:"urgency_load"`
	DeletedGoals []struct {
		ID string `json:"id"`
	} `json:"deleted_goals,omitempty"`
}

// Goal represents a Beeminder goal
type Goal struct {
	Slug            string          `json:"slug"`
	Title           string          `json:"title"`
	Description     interface{}     `json:"description"`
	GoalType        string          `json:"goal_type"`
	Autodata        string          `json:"autodata"`
	AutodataConfig  interface{}     `json:"autodata_config"`
	HealthkitMetric string          `json:"healthkitmetric"`
	Deadline        int             `json:"deadline"`
	Leadtime        int             `json:"leadtime"`
	Alertstart      int             `json:"alertstart"`
	Pledge          float64         `json:"pledge"`
	PledgeCap       float64         `json:"pledge_cap"`
	Rosy            bool            `json:"rosy"`
	Yaw             float64         `json:"yaw"`
	Rate            float64         `json:"rate"`
	Gunits          string          `json:"gunits"`
	Runits          string          `json:"runits"`
	Limsum          string          `json:"limsum"`
	Frozen          bool            `json:"frozen"`
	Lost            bool            `json:"lost"`
	Won             bool            `json:"won"`
	Contract        Contract        `json:"contract"`
	Delta           float64         `json:"delta"`
	DeltaText       string          `json:"delta_text"`
	SafeBuf         int64           `json:"safebuf"`
	SafeBump        float64         `json:"safebump"`
	Curval          float64         `json:"curval"`
	Lastday         int64           `json:"lastday"`
	Todayta         bool            `json:"todayta"`
	Hhmmformat      bool            `json:"hhmmformat"`
	Kyoom           bool            `json:"kyoom"`
	Odom            bool            `json:"odom"`
	Aggday          string          `json:"aggday"`
	Plotall         bool            `json:"plotall"`
	Steppy          bool            `json:"steppy"`
	Integery        bool            `json:"integery"`
	Graphsum        string          `json:"graphsum"`
	Timezone        string          `json:"timezone"`
	Losedate        int64           `json:"losedate"`
	Road            [][]interface{} `json:"road"`
	Roadall         [][]interface{} `json:"roadall"`
	Curday          float64         `json:"curday"`
	LastDatapoint   Datapoint       `json:"last_datapoint"`
	Datapoints      []Datapoint     `json:"datapoints,omitempty"`
}

// Contract represents the contract details of a goal
type Contract struct {
	Amount        float64     `json:"amount"`
	PendingAmount interface{} `json:"pending_amount"`
	PendingAt     interface{} `json:"pending_at"`
	StepdownAt    interface{} `json:"stepdown_at"`
}

// Datapoint represents a Beeminder datapoint
type Datapoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
	Comment   string  `json:"comment"`
	ID        string  `json:"id"`
	UpdatedAt int64   `json:"updated_at"`
	Daystamp  string  `json:"daystamp"`
	Origin    string  `json:"origin"`
	Canonical string  `json:"canonical"`
	Fulltext  string  `json:"fulltext"`
	RequestID string  `json:"requestid"`
	CreatedAt string  `json:"created_at"`
	Creator   string  `json:"creator"`
	Urtext    string  `json:"urtext"`
}

// GetUserParams represents the parameters for the GetUser method
type GetUserParams struct {
	Username        string
	Associations    bool
	DiffSince       *int64
	Skinny          bool
	DatapointsCount *int64
}

// ============================== //
// ========= NOTION API ========= //
// ============================== //
