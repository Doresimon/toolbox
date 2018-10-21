package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckQuadrcticResidues(t *testing.T) {
	assert := assert.New(t)

	a := 9
	p := 11
	r := CheckQuadrcticResidues(a, p)

	assert.Equal(r, true, "The result shuold be true.")

	a = 10
	p = 11
	r = CheckQuadrcticResidues(a, p)

	assert.Equal(r, false, "The result shuold be false.")
}
