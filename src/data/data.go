package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MydroX/short-circuit/pkg/geojson"
	"github.com/MydroX/short-circuit/src/models"
	"github.com/google/uuid"

	geojsonModels "github.com/MydroX/short-circuit/pkg/geojson/models"
)

func Load() {
	// TODO: find a way to load certain dataset or not

	if os.Getenv("GEOJSON_FRANCE") != "" {
		polygons := GetPolygonCFromFeatureCollection(os.Getenv("GEOJSON_FRANCE"))
		fmt.Println(len(polygons))
		return
	}

	if os.Getenv("GEOJSON_TEST") != "" {
		_ = GetPolygonCFromFeatureCollection(os.Getenv("GEOJSON_TEST"))
	}

	log.Fatalf("unable to load any dataset")
}

func GetPolygonCFromFeatureCollection(fileName string) []*models.PolygonC {
	jsonByte := geojson.Read(fileName)

	var featureCollection geojsonModels.FeatureCollection
	err := json.Unmarshal(jsonByte, &featureCollection)
	if err != nil {
		log.Fatalf("failed to unmarshal %v: %v", fileName, err)
	}

	var polygonsC []*models.PolygonC

	for _, f := range featureCollection.Features {
		if f.Geometry.Type == "MultiPolygon" {
			var geometry [][][][]float32
			json.Unmarshal(f.Geometry.GeometryData, &geometry)

			for _, multipolygon := range geometry {
				if len(multipolygon) > 1 {
					log.Fatalf("polygon holes support is not implemented")
				}

				for _, polygon := range multipolygon {

					polygonC := polygonGeoToPolygonC(polygon)
					polygonC.Uuid = uuid.NewString()

					polygonsC = append(polygonsC, &polygonC)
				}
			}

			polygonsMap := make(map[string]bool)
			for _, p := range polygonsC {
				polygonsMap[p.Uuid] = true
			}

			for _, p := range polygonsC {
				var linkedUuid []string

				for pUuid := range polygonsMap {
					if pUuid != p.Uuid {
						linkedUuid = append(linkedUuid, pUuid)
					}
				}
				p.LinkedPolygons = linkedUuid
			}
		}

		if f.Geometry.Type == "Polygon" {
			var geometry [][][]float32
			json.Unmarshal(f.Geometry.GeometryData, &geometry)

			var mainPolygon models.PolygonC

			for i, polygon := range geometry {

				var polygonC models.PolygonC

				if i == 0 {
					polygonC = polygonGeoToPolygonC(polygon)
					mainPolygon.Uuid = polygonC.Uuid
				} else {
					if i != 0 {
						polygonC = polygonGeoToPolygonC(polygon)
						mainPolygon.Holes = append(mainPolygon.Holes, polygonC.Uuid)
					}
				}

				polygonsC = append(polygonsC, &polygonC)
			}
		}
	}

	return polygonsC
}

func polygonGeoToPolygonC(polygon [][]float32) (polygonC models.PolygonC) {
	for _, point := range polygon {

		p := make([]float32, 2)

		p[0] = point[0]
		p[1] = point[1]

		polygonC.Polygon.Points = append(polygonC.Polygon.Points, p)
	}
	polygonC.Uuid = uuid.NewString()

	return polygonC
}
