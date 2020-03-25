package models

type RRPtr struct {
	Value int `json:"value"`

	DisableFlag string `json:"disableflag,omitempty"`
}

type PtrAttributes struct {
	Name       string        `json:"name,omitempty"`
	TTL        int           `json:"ttl"`
	NoAnswer   string        `json:"noanswer,omitempty"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdregion,omitempty"`
	Type       string        `json:"type,omitempty"`
	ParentID   int           `json:"parentid,omitempty"`
	Parent     string        `json:"parent,omitempty"`
	Source     string        `json:"source,omitempty"`
	RoundRobin []interface{} `json:"roundRobin"`
}
