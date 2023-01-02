package geojson

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Read(fileName string) []byte {
	jsonFile, err := os.Open(fmt.Sprintf("data/%v", fileName))
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer jsonFile.Close()

	byteJson, _ := io.ReadAll(jsonFile)

	return byteJson
}
