package models

type TemplateAttributes struct {
	Name          []string `json:"name"`
	Domain        int      `json:"Domain,omitempty"`
	HasGtdRegions bool     `json:"hasGtdRegions,omitempty"`
	HasGeoIP      bool     `json:"hasGeoIP,omitempty"`
	Version       int      `json:"version,omitempty"`
}
