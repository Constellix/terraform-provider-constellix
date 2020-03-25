package models

type DNSAttributes struct {
	Name         string        `json:"name,omitempty"`
	FQDN         string        `json:"fqdn,omitempty"`
	Resolver     string        `json:"resolver,omitempty"`
	Host         string        `json:"host,omitempty"`
	Port         int           `json:"port,omitempty"`
	ProtocolType string        `json:"protocolType,omitempty"`
	CheckSites   []interface{} `json:"checkSites,omitempty"`
}
