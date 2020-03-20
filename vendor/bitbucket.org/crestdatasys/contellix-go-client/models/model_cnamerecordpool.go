package models

type ValuesCname struct {
	Value       string `json:"value"`
	Weight      int    `json:"weight,omitempty"`
	DisableFlag bool   `json:"disableFlag,omitempty"`
	CheckId     int    `json:"checkId,omitempty"`
	Policy      string `json:"policy"`
}

type CnameRecordPoolAttributes struct {
	Name                 string `json:"name"`
	NumReturn            int    `json:"numReturn"`
	MinAvailableFailover int    `json:"minAvailableFailover"`
	FailedFlag           string `json:"failedFlag,omitempty"`
	DisableFlag1         string `json:"disableFlag,omitempty"`
	Note                 string `json:"note,omitempty"`
	Version              int    `json:"version,omitempty"`

	ValuesCname []interface{} `json:"values"`
}
