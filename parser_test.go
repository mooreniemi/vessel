package vessel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parsedIntMap = [][]int{
	//    0  1
	[]int{0, 0, 1, 1, 1},
	//       2  3  4
	[]int{1, 0, 0, 0, 1},
	//       5
	[]int{1, 0, 1, 1, 1},
	//       6  7
	[]int{1, 0, 0, 1, 1},
	[]int{1, 1, 0, 1, 1}}

var parsedMap = [][]string{
	[]string{"0", "0", "1", "1", "1"},
	[]string{"1", "0", "0", "0", "1"},
	[]string{"1", "0", "1", "1", "1"},
	[]string{"1", "0", "0", "1", "1"},
	[]string{"1", "1", "0", "1", "1"}}

func TestParseVesselMap(t *testing.T) {
	assert.Equal(t, parsedMap, ParseVesselMap(), "")
}

func TestByCoordinates(t *testing.T) {
	x := 1
	y := 2
	assert.Equal(t, parsedIntMap[x][y], VesselMapAsInts(ParseVesselMap())[x][y], "")
	assert.Equal(t, 0, VesselMapAsInts(ParseVesselMap())[x][y], "")
}
