package models

type GeolocationArecord struct {
	GeoIpUserRegion []int `json:"geoipUserRegion,omitempty"`
	Drop            bool  `json:"drop,omitempty"`
	GeoIpProximity  int   `json:"geoipProximity,omitempty"`
	GeoIpFailOver   bool  `json:"geoipFailover,omitempty"`
}

type RoundrobinArecord struct {
	Value       string `json:"value,omitempty"`
	DisableFlag string `json:"disableFlag,omitempty"`
}

type ValuesRCDFArecord struct {
	Value         string `json:"value,omitempty"`
	SortOrderRCDF int    `json:"sortOrder,omitempty"`
	CheckId       int    `json:"checkid,omitempty"`
	DisableFlag   bool   `json:"disableFlag,omitempty"`
}

type RRFAARecord struct {
	Value           string `json:"value,omitempty"`
	SortOrder       int    `json:"sortOrder,omitempty"`
	DisableFlagRRFA bool   `json:"disableFlag,omitempty"`
}

type RCDFAARecord struct {
	Values            []interface{} `json:"values,omitempty"`
	DisableFlagRCDFA  bool          `json:"disabled,omitempty"`
	FailoverTypeRCDFA int           `json:"failoverType,omitempty"`
}

type ARecordAttributes struct {
	//DomainId           string             `json: "domainid,omitempty"`
	Name                string              `json:"name,omitempty"`
	TTL                 int                 `json:"ttl"`
	GeoLocation         *GeolocationArecord `json:"geolocation,omitempty"`
	RecordOption        string              `json:"recordOption,omitempty"`
	NoAnswer            bool                `json:"noAnswer,"`
	Note                string              `json:"note,omitempty"`
	GtdRegion           int                 `json:"gtdRegion,omitempty"`
	Type                string              `json:"type,omitempty"`
	ContactId           []int               `json:"contactIds,omitempty"`
	RoundRobin          []interface{}       `json:"roundRobin,omitempty"`
	Pools               []int               `json:"pools,omitempty"`
	RoundRobinFailoverA []interface{}       `json:"roundRobinFailover,omitempty"`
	RecordFailoverA     *RCDFAARecord       `json:"recordFailover,omitempty"`
}
