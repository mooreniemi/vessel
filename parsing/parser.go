package parser

import (
	"fmt"
	v "github.com/mooreniemi/vessel/vessel"
	"io/ioutil"
	"log"
)

// ParseVesselYaml expects a resources directory
// with a yaml file to turn into Chambers
// http://www.yamllint.com/ also helpful
func ParseVesselYaml() (v.Vessel, error) {
	data, err := ioutil.ReadFile("resources/vessel.yml")

	if err != nil {
		log.Fatal(err)
	}

	var vessel v.Vessel
	if err := vessel.Parse(data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("root chamber: %+v\n", vessel.Chambers[0])

	return vessel, err
}
