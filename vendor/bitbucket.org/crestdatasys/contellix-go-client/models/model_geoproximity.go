package models

type GeoProximityAttributes struct {
	Name                string  `json:"name,omitempty"`
	Country             string  `json:"country,omitempty"`
	RegionStateProvince string  `json:"region,omitempty"`
	City                int     `json:"city,omitempty"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
}
