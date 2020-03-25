package models

type RRSpf struct {
	Value string `json:"value"`

	DisableFlag bool `json:"disableflag,omitempty"`
}

type SpfAttributes struct {
	Name       string        `json:"name,omitempty"`
	TTL        int           `json:"ttl"`
	NoAnswer   bool          `json:"noanswer,omitempty"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdregion,omitempty"`
	Type       string        `json:"type,omitempty"`
	ParentID   int           `json:"parentid,omitempty"`
	Parent     string        `json:"parent,omitempty"`
	Source     string        `json:"source,omitempty"`
	RoundRobin []interface{} `json:"roundRobin"`
}
