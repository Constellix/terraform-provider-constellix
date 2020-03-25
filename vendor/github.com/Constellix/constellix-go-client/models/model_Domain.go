package models

type Soa struct {
	PrimaryNameServer string `json:"primaryNameserver,omitempty"`
	Email             string `json:"email,omitempty"`
	TTL               int    `json:"ttl,omitempty"`
	Serial            int    `json:"serial,omitempty"`
	Refresh           int    `json:"refresh,omitempty"`
	Retry             int    `json:"retry,omitempty"`
	Expire            int    `json:"expire,omitempty"`
	NegCache          int    `json:"negCache,omitempty"`
}

type DomainAttributes struct {
	Name            []string      `json:"names"`
	TypeID          int           `json:"typeId,omitempty"`
	HasGtdRegions   bool          `json:"hasGtdRegions,omitempty"`
	HasGeoIP        bool          `json:"hasGeoIP,omitempty"`
	NameserverGroup string        `json:"nameserverGroup,omitempty"`
	Nameservers     []string      `json:"nameservers,omitempty"`
	Note            string        `json:"note,omitempty"`
	Version         int           `json:"version,omitempty"`
	Status          string        `json:"status,omitempty"`
	Tags            []interface{} `json:"tags,omitempty"`

	Soa *Soa `json:"soa,omitempty"`
}
