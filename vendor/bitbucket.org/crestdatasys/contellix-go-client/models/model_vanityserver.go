package models

type VanityNameserverAttributes struct {
	Name                 string `json:"name,omitempty"`
	IsDefault            bool   `json:"isDefault,omitempty"`
	IsPublic             bool   `json:"isPublic,omitempty"`
	NameserverGroup      int    `json:"nameserverGroup,omitempty"`
	NameserverGroupName  string `json:"nameserverGroupName,omitempty"`
	NameserverListString string `json:"nameserversListString,omitempty"`
}
