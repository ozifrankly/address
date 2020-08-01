package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEP(t *testing.T) {
	addr, err := fetch("01001000")
	assert.Equal(t, "São Paulo", addr.City)
	assert.Equal(t, "Sé", addr.District)
	assert.NotEqual(t, "", addr.Complement)
	assert.NotNil(t, addr.Complement)
	assert.NotEqual(t, "", addr.Street)
	assert.NotNil(t, addr.Street)
	assert.Nil(t, err)
}

func TestGetInvalidCEP(t *testing.T) {
	addr, err := fetch("00000000")
	assert.Nil(t, addr)
	assert.Error(t, err)
}
