package models

type NSAttributes struct {
	Name       string        `json:"name,omitempty"`
	Ttl        int           `json:"ttl"`
	NoAnswer   bool          `json:"noAnswer,omitempty"`
	RoundRobin []interface{} `json:"roundRobin"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdRegion,omitempty"`
	Type       string        `json:"type,omitempty"`
}

type RoundRobinNS struct {
	Value       string `json:"value,omitempty"`
	DisableFlag bool   `json:"disableFlag,omitempty"`
}
