package models

type Polygon struct {
	Points     [][]float32
	Properties Properties
}

type Properties struct {
	Code              string
	PopulationDensity uint16
}

type PolygonC struct {
	Uuid           string
	Polygon        Polygon
	Holes          []string
	LinkedPolygons []string
	Properties     Properties
}
