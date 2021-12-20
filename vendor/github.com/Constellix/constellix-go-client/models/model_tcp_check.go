package models

type TCPCheckAttributes struct {
	Name                      string        `json:"name,omitempty"`
	Host                      string        `json:"host"`
	Port                      int           `json:"port"`
	Ipversion                 string        `json:"ipVersion,omitempty"`
	Checksites                []interface{} `json:"checkSites"`
	Interval                  string        `json:"interval,omitempty"`
	IntervalPolicy            string        `json:"monitorIntervalPolicy,omitempty"`
	VerificationPolicy        string        `json:"verificationPolicy,omitempty"`
	StringToSend              string        `json:"stringToSend,omitempty"`
	StringToReceive           string        `json:"stringToReceive,omitempty"`
	NotificationGroups        []int         `json:"notificationGroups,omitempty"`
	NotificationReportTimeout int           `json:"notificationReportTimeout,omitempty"`
}
