package models

type HTTPRedirectionAttributes struct {
	Name           string `json:"name,omitempty"`
	TTL            int    `json:"ttl"`
	NoAnswer       bool   `json:"noAnswer,omitempty"`
	Title          string `json:"title,omitempty"`
	Keywords       string `json:"keywords,omitempty"`
	Description    string `json:"description,omitempty"`
	Note           string `json:"note,omitempty"`
	GtdRegion      int    `json:"gtdRegion,omitempty"`
	URL            string `json:"url"`
	Type           string `json:"type,omitempty"`
	Hardlinkflag   bool   `json:"hardlinkFlag,omitempty"`
	RedirectTypeID int    `json:"redirectTypeId"`
	ParentID       int    `json:"parentId,omitempty"`
	Parent         string `json:"parent,omitempty"`
	Source         string `json:"source,omitempty"`
}
