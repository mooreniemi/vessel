package vesselparser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Chamber is a node in the vessel
type Chamber struct {
	Desc     string
	Doors    []int
	DoorDesc string `yaml:"doorDesc"`
	ID       int
	Items    []int
}

// Vessel is a graph of chambers
type Vessel struct {
	Chambers []*Chamber
}

// Parse yaml into Vessel struct
func (c *Vessel) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// ParseVesselYaml expects a resources directory
// with a yaml file to turn into Chambers
// http://www.yamllint.com/ also helpful
func ParseVesselYaml() (Vessel, error) {
	data, err := ioutil.ReadFile("resources/vessel.yml")

	if err != nil {
		log.Fatal(err)
	}

	var vessel Vessel
	if err := vessel.Parse(data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("root chamber: %+v\n", vessel.Chambers[0])

	return vessel, err
}
