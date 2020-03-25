package models

type TCPCheckAttributes struct {
	Name       string        `json:"name,omitempty"`
	Host       string        `json:"host"`
	Port       int           `json:"port"`
	Ipversion  string        `json:"ipVersion,omitempty"`
	Checksites []interface{} `json:"checkSites"`
}
