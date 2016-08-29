package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parsedMap = [][]int{
	[]int{0, 0, 1, 1, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 1, 1, 1},
	[]int{1, 0, 0, 1, 1},
	[]int{1, 1, 1, 1, 1}}

func TestParseVesselMap(t *testing.T) {
	assert.Equal(t, parsedMap, ParseVesselMap(), "")
}

func TestByCoordinates(t *testing.T) {
	x := 1
	y := 2
	assert.Equal(t, parsedMap[x][y], ParseVesselMap()[x][y], "")
	assert.Equal(t, 0, ParseVesselMap()[x][y], "")
}
