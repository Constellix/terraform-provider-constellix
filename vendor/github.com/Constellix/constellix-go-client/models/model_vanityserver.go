package models

type VanityNameserverAttributes struct {
	Name                 string `json:"name,omitempty"`
	IsDefault            bool   `json:"isDefault"`
	IsPublic             bool   `json:"isPublic"`
	NameserverGroup      int    `json:"nameserverGroup,omitempty"`
	NameserverGroupName  string `json:"nameserverGroupName,omitempty"`
	NameserverListString string `json:"nameserversListString,omitempty"`
}
