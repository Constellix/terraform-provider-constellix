package models

type NAPTRAttributes struct {
	Name       string        `json:"name,omitempty"`
	Ttl        int           `json:"ttl"`
	NoAnswer   bool          `json:"noAnswer"`
	RoundRobin []interface{} `json:"roundRobin"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdRegion,omitempty"`
	Type       string        `json:"type,omitempty"`
}

type RoundRobinNaptr struct {
	Order             int    `json:"order,omitempty"`
	Preference        int    `json:"preference,omitempty"`
	Flags             string `json:"flags,omitempty"`
	Service           string `json:"service,omitempty"`
	RegularExpression string `json:"regularExpression,omitempty"`
	Replacement       string `json:"replacement,omitempty"`
	DisableFlag       bool   `json:"disableFlag,omitempty"`
}
