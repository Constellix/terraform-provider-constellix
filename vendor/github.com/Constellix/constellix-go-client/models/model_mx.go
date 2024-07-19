package models

type RoundR struct {
	Value       string `json:"value"`
	Level       string `json:"level"`
	DisableFlag bool   `json:"disableflag"`
}

type MXAttributes struct {
	Name       string        `json:"name,omitempty"`
	TTL        int           `json:"ttl"`
	NoAnswer   bool          `json:"noanswer"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdregion,omitempty"`
	Type       string        `json:"type,omitempty"`
	ParentId   int           `json:"parentid,omitempty"`
	Parent     string        `json:"parent,omitempty"`
	Source     string        `json:"source,omitempty"`
	RoundRobin []interface{} `json:"roundRobin"`
}
