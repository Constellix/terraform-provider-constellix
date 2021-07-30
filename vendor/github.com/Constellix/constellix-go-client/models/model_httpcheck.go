package models

type HttpcheckAttr struct {
	Name               string        `json:"name,omitempty"`
	Host               string        `json:"host,omitempty"`
	Ipversion          string        `json:"ipVersion,omitempty"`
	Port               int           `json:"port"`
	ProtoType          string        `json:"protocolType"`
	Checksites         []interface{} `json:"checkSites"`
	Interval           string        `json:"interval,omitempty"`
	IntervalPolicy     string        `json:"monitorIntervalPolicy,omitempty"`
	VerificationPolicy string        `json:"verificationPolicy,omitempty"`
	FQDN               string        `json:"fqdn,omitempty"`
	PATH               string        `json:"path,omitempty"`
	SearchString       string        `json:"searchString,omitempty"`
	ExpectedStatus     int           `json:"expectedStatusCode,omitempty"`
	NotificationGroups []int         `json:"notificationGroups,omitempty"`
}
