package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVesselMap(t *testing.T) {
	parsedMap := [][]int{
		[]int{0, 0, 1, 1, 1},
		[]int{1, 0, 0, 0, 1},
		[]int{1, 0, 1, 1, 1},
		[]int{1, 0, 0, 1, 1},
		[]int{1, 1, 1, 1, 1}}
	assert.Equal(t, parsedMap, ParseVesselMap(), "")
}
