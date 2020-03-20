package models

type ValuesAAAApool struct {
	Value       string `json:"value,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	Disableflag bool   `json:"disableFlag,omitempty"`
	Checkid     int    `json:"checkId,omitempty"`
	Policy      string `json:"policy,omitempty"`
}

type AAAArecordPoolAttributes struct {
	Name             string        `json:"name,omitempty"`
	NumReturn        int           `json:"numReturn,omitempty"`
	MinavailFailover int           `json:"minAvailableFailover,omitempty"`
	Failedflag       bool          `json:"failedFlag,omitempty"`
	Disableflag      bool          `json:"disableFlag,omitempty"`
	Values           []interface{} `json:"values,omitempty"`
	Note             string        `json:"note,omitempty"`
}
