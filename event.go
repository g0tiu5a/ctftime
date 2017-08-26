package ctftime

import "time"

/* Struct */

type Event struct {
	Organizers    []Organizer `json:"organizers"`
	OnSite        bool        `json:"onsite"`
	Finish        time.Time   `json:"finish"`
	Description   string      `json:"description"`
	Weight        float64     `json:"weight"`
	Title         string      `json:"title"`
	Url           string      `json:"url"`
	IsVotableNow  bool        `json:"is_votable_now"`
	Restrictions  string      `json:"restrictions"`
	Format        string      `json:"format"`
	Start         time.Time   `json:"start"`
	Participants  int         `json:"participants"`
	CtftimeUrl    string      `json:"ctftime_url"`
	Location      string      `json:"location"`
	LiveFeed      string      `json:"live_feed"`
	PublicVotable bool        `json:"public_votable"`
	Duration      Duration    `json:"duration"`
	Logo          string      `json:"logo"`
	FormatId      int         `json:"format_id"`
	Id            int         `json:"id"`
	CtfId         int         `json:"ctf_id"`
}

type Organizer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Duration struct {
	Hours int `json:"hours"`
	Days  int `json:"days"`
}
