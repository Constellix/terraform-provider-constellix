package models

type SRVAttributes struct {
	Name       string        `json:"name,omitempty"`
	TTL        int           `json:"ttl"`
	NoAnswer   bool          `json:"noAnswer,omitempty"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdRegion,omitempty"`
	Type       string        `json:"type,omitempty"`
	RoundRobin []interface{} `json:"roundRobin,omitempty"`
}

type RoundRobinSRV struct {
	Value       string `json:"value,omitempty"`
	Port        int    `json:"port,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	DisableFlag bool   `json:"disableflag,omitempty"`
}
