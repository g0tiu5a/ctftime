package ctftime

type Team struct {
	TeamId   int64   `json:"team_id"`
	TeamName string  `json:"team_name"`
	Points   float64 `json:"points"`
}

type Top10 map[string][]Team
