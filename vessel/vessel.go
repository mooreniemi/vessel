package vessel

import (
	"gopkg.in/yaml.v2"
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
