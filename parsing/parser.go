package parser

import (
	"bytes"
	"encoding/csv"
	bd "github.com/mooreniemi/vessel/bindata"
	v "github.com/mooreniemi/vessel/vessel"
	"log"
	"strconv"
)

// ParseVesselYaml expects a resources directory
// with a yaml file to turn into Chambers
// http://www.yamllint.com/ also helpful
func ParseVesselYaml() (v.Vessel, error) {
	data, err := bd.Asset("data/vessel.yml")

	if err != nil {
		log.Fatal(err)
	}

	var vessel v.Vessel
	if err := vessel.Parse(data); err != nil {
		log.Fatal(err)
	}

	return vessel, err
}

// ParseVesselMap expects a resources direcotry
// with a csv file to turn into a map
func ParseVesselMap() [][]string {
	data, err := bd.Asset("data/vessel.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(data))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

// VesselMapAsInts gives back the map in [][]int
func VesselMapAsInts(stringMap [][]string) [][]int {
	// TODO change to use https://github.com/gocarina/gocsv ?
	intMap := make([][]int, len(stringMap))
	for i, row := range stringMap {
		intMap[i] = make([]int, len(row))
		for j := range row {
			intVal, _ := strconv.Atoi(stringMap[i][j])
			intMap[i][j] = intVal
		}
	}
	return intMap
}
