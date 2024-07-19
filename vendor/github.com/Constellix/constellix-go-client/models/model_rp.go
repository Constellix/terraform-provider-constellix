package models

type RoundRobinRp struct {
	Mailbox     string `json:"mailbox,omitempty"`
	Txt         string `json:"txt,omitempty"`
	DisableFlag string `json:"disableFlag,omitempty"`
}

type RPAttributes struct {
	Name       string        `json:"name,omitempty"`
	TTL        int           `json:"ttl,omitempty"`
	NoAnswer   bool          `json:"noanswer"`
	Note       string        `json:"note,omitempty"`
	GtdRegion  int           `json:"gtdregion,omitempty"`
	Type       string        `json:"type,omitempty"`
	ParentId   int           `json:"parentid,omitempty"`
	Parent     string        `json:"parent,omitempty"`
	Source     string        `json:"source,omitempty"`
	RoundRobin []interface{} `json:"roundRobin"`
}
