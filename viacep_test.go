package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEP(t *testing.T) {
	address, err := fetchViacep("01001000")
	assert.Equal(t, "São Paulo", address.City)
	assert.Equal(t, "Sé", address.District)
	assert.Equal(t, "lado ímpar", address.Complement)
	assert.Equal(t, "Praça da Sé", address.Street)
	assert.Nil(t, err)
}

func TestGetInvalidCEP(t *testing.T) {
	address, err := fetchViacep("00000000")
	assert.Nil(t, address)
	assert.Error(t, err)
}
