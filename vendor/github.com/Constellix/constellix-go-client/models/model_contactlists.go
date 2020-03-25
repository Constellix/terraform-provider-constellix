package models

type ContactListAttributes struct {
	Name string `json:"name"`

	EmailAddresses []string `json:"emailAddresses"`
	//ContactEmailsCount int      `json:"con,omitempty"`
	//ActiveFlag         bool     `json:"activeFlag,omitempty"`
}
