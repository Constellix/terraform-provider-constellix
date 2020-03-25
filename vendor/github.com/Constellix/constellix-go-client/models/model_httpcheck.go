package models

type HttpcheckAttr struct {
	Name       string        `json:"name,omitempty"`
	Host       string        `json:"host,omitempty"`
	Ipversion  string        `json:"ipVersion,omitempty"`
	Port       int           `json:"port"`
	ProtoType  string        `json:"protocolType"`
	Checksites []interface{} `json:"checkSites"`
}
