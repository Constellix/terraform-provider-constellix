package models

type RRCert struct {
	CertificateType int    `json:"certificateType"`
	KeyTag          int    `json:"keyTag"`
	DisableFlag     bool   `json:"disableflag,omitempty"`
	Algorithm       int    `json:"algorithm"`
	Certificate     string `json:"certificate"`
}

type CertAttributes struct {
	Name       string   `json:"name,omitempty"`
	TTL        int      `json:"ttl"`
	NoAnswer   bool     `json:"noanswer"`
	Note       string   `json:"note,omitempty"`
	GtdRegion  int      `json:"gtdregion,omitempty"`
	Type       string   `json:"type,omitempty"`
	ParentID   int      `json:"parentid,omitempty"`
	Parent     string   `json:"parent,omitempty"`
	Source     string   `json:"source,omitempty"`
	RoundRobin []interface{} `json:"roundRobin"`
}
