package ctftime

type Team struct {
	TeamId   int64   `json:"team_id"`
	TeamName string  `json:"team_name"`
	Points   float64 `json:"points"`
}

type Top10 []Team
type Top10s map[string]Top10
