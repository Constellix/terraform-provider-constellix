package models

type RoundRobinCaa struct {
	CaaproviderId int    `json:"caaProviderId,omitempty"`
	Tag           string `json:"tag,omitempty"`
	Data          string `json:"data,omitempty"`
	Flag          string `json:"flag,omitempty"`
	DisableFlag   bool   `json:"disableflag"`
}

type CaaAttributes struct {
	Name       string        `json:"name,omitempty"`
	TTL        int           `json:"ttl,omitempty"`
	NoAnswer   bool          `json:"noanswer"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdregion,omitempty"`
	Type       string        `json:"type,omitempty"`
	ParentId   int           `json:"parentid,omitempty"`
	Parent     string        `json:"parent,omitempty"`
	Source     string        `json:"source,omitempty"`
	RoundRobin []interface{} `json:"roundRobin,omitempty"`
}
