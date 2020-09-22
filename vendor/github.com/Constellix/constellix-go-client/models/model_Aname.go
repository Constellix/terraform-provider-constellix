package models

type AnameAttributes struct {
	//Id              int      `json:"id"`
	//DomainID       string   `json:"domainid"`
	Name                string            `json:"name,omitempty"`
	TTL                 int               `json:"ttl"`
	RecordOption        string            `json:"recordOption,omitempty"`
	NoAnswer            bool              `json:"noAnswer,omitempty"`
	Note                string            `json:"note,omitempty"`
	GtdRegion           int               `json:"gtdRegion,omitempty"`
	Type                string            `json:"type,omitempty"`
	ContactIDs          []int             `json:"contactids,omitempty"`
	RoundRobin          []interface{}     `json:"roundRobin,omitempty"`
	Pools               []int             `json:"pools,omitempty"`
	GeoLocation         *GeolocationANAME `json:"geolocation,omitempty"`
	RecordFailoverAname *RCDFAname        `json:"recordFailover,omitempty"`
}

type GeolocationANAME struct {
	GeoIpUserRegion []int `json:"geoipUserRegion,omitempty"`
	Drop            bool  `json:"drop,omitempty"`
	GeoIpProximity  int   `json:"geoipProximity,omitempty"`
	GeoIpFailOver   bool  `json:"geoipFailover,omitempty"`
}

type AnameRoundRobin struct {
	Value       string `json:"value,omitempty"`
	DisableFlag string `json:"disableFlag,omitempty"`
}

type RCDFAname struct {
	Values            []interface{} `json:"values,omitempty"`
	FailoverTypeAname int           `json:"failoverType,omitempty"`
	DisableFlagAname  bool          `json:"disabled,omitempty"`
}

type ValuesAname struct {
	Value         string `json:"value,omitempty"`
	SortOrderRCDF int    `json:"sortOrder,omitempty"`
	DisableFlag   bool   `json:"disableFlag,omitempty"`
	CheckID       int    `json:"checkid,omitempty"`
}
