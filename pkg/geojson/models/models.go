package models

import "encoding/json"

type FeatureCollection struct {
	Type       string     `json:"type"`
	Features   []Feature  `json:"features"`
	Properties Properties `json:"properties"`
}

type Feature struct {
	Type     string   `json:"type"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type         string          `json:"type"`
	GeometryData json.RawMessage `json:"coordinates"`
}

type Properties struct {
	Code     string `json:"code"`
	Cityname string `json:"nom"`
}
