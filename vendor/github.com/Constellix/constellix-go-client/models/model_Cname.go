package models

type GeolocationCrecord struct {
	GeoIpUserRegion []int `json:"geoipUserRegion,omitempty"`
	Drop            bool  `json:"drop,omitempty"`
	GeoIpProximity  int   `json:"geoipProximity,omitempty"`
	GeoIpFailOver   bool  `json:"geoipFailover,omitempty"`
}

type ValuesRCDFCrecord struct {
	Value         string `json:"value,omitempty"`
	SortOrderRCDF int    `json:"sortOrder,omitempty"`
	CheckId       int    `json:"checkid,omitempty"`
	DisableFlag   bool   `json:"disableFlag,omitempty"`
}

type RCDFACRecord struct {
	Values            []interface{} `json:"values,omitempty"`
	DisableFlagRCDFA  bool          `json:"disabled,omitempty"`
	FailoverTypeRCDFA int           `json:"failoverType,omitempty"`
}

type CRecordAttributes struct {
	//DomainId           string             `json: "domainid,omitempty"`
	Name            string              `json:"name,omitempty"`
	Host            string              `json:"host,omitempty"`
	TTL             int                 `json:"ttl"`
	GeoLocation     *GeolocationCrecord `json:"geolocation,omitempty"`
	RecordOption    string              `json:"recordOption,omitempty"`
	NoAnswer        bool                `json:"noAnswer,omitempty"`
	Note            string              `json:"note,omitempty"`
	GtdRegion       int                 `json:"gtdRegion,omitempty"`
	Type            string              `json:"type,omitempty"`
	ContactId       []int               `json:"contactId,omitempty"`
	Pools           []int               `json:"pools,omitempty"`
	RecordFailoverA *RCDFACRecord       `json:"recordFailover,omitempty"`
}
