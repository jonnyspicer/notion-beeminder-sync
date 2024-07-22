package unified

// ============================== //
// =========== UNIFIED ========== //
// ============================== //

// An intermediary type that represents a Beeminder goal in Notion
type Goal struct {
	Goal      string   `json:"goal"`
	Deadline  int64    `json:"deadline"` // Unix timestamp
	Delta     float64  `json:"delta"`
	GoalUnits string   `json:"goal_units"`
	SafeDays  int64    `json:"safe_days"`
	Pledge    float64  `json:"pledge"`
	Link      string   `json:"link"`
	Autodata  string   `json:"autodata"`
	GoalType  GoalType `json:"goal_type"`
	RateUnit  RateUnit `json:"rate_unit"`
}

type GoalType string

type RateUnit string

const (
	DoMore     GoalType = "do_more"
	Odometer   GoalType = "odometer"
	WeightLoss GoalType = "weight_loss"
	WeightGain GoalType = "weight_gain"
	InboxFewer GoalType = "inbox_fewer"
	DoLess     GoalType = "do_less"
	Custom     GoalType = "custom"
)

const (
	Hourly  RateUnit = "hourly"
	Daily   RateUnit = "daily"
	Weekly  RateUnit = "weekly"
	Monthly RateUnit = "monthly"
	Yearly  RateUnit = "yearly"
)
