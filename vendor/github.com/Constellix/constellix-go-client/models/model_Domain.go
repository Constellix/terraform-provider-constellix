package models

type Soa struct {
	PrimaryNameServer string `json:"primaryNameserver,omitempty"`
	Email             string `json:"email,omitempty"`
	TTL               string `json:"ttl,omitempty"`
	Serial            string `json:"serial,omitempty"`
	Refresh           string `json:"refresh,omitempty"`
	Retry             string `json:"retry,omitempty"`
	Expire            string `json:"expire,omitempty"`
	NegCache          string `json:"negCache,omitempty"`
}

type DomainAttributes struct {
	Name             []string      `json:"names,omitempty"`
	TypeID           int           `json:"typeId,omitempty"`
	HasGtdRegions    bool          `json:"hasGtdRegions"`
	HasGeoIP         bool          `json:"hasGeoIP"`
	NameserverGroup  string        `json:"nameserverGroup,omitempty"`
	VanityNameServer string        `json:"vanityNameServer,omitempty"`
	Nameservers      []string      `json:"nameservers,omitempty"`
	Note             string        `json:"note,omitempty"`
	Version          int           `json:"version,omitempty"`
	Status           string        `json:"status,omitempty"`
	Tags             []interface{} `json:"tags,omitempty"`
	Template         int           `json:"template,omitempty"`
	Soa              *Soa          `json:"soa,omitempty"`
}
