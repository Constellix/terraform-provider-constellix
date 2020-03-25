package models

type IPAddresses struct {
	IPV4Addresses []IPV4Addresses `json:"ipv4Addresses,omitempty"`
	IPV6Addresses []IPV6Addresses `json:"ipv6Addresses,omitempty"`
}

type IPV4Addresses struct {
	IPV4 string `json:"ipv4:omitempty"`
}

type IPV6Addresses struct {
	IPV6 string `json:"ipv6:omitempty"`
}

type IPFilterAttributes struct {
	// Id              int      `json:"id,omitempty"`
	Name             string        `json:"name"`
	GeoIPContinents  []string      `json:"geoipContinents,omitempty"`
	GeoIPRegions     []string      `json:"geoipRegions,omitempty"`
	GeoIPCountries   []string      `json:"geoipCountries,omitempty"`
	Asn              []int         `json:"asn,omitempty"`
	FilterRulesLimit int           `json:"filterRulesLimit,omitempty"`
	IPAddresses      []interface{} `json:"ipaddrs,omitempty"`
}
