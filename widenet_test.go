package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEPFromWidenet(t *testing.T) {
	address, err := fetchWidenet("01001000")
	assert.Equal(t, "São Paulo", address.City)
	assert.Equal(t, "Sé", address.District)
	assert.Equal(t, "", address.Complement)
	assert.Equal(t, "Praça da Sé - lado ímpar", address.Street)
	assert.Nil(t, err)
}

func TestGetInvalidCEPFromWidenet(t *testing.T) {
	address, err := fetchWidenet("00000000")
	assert.Nil(t, address)
	assert.Error(t, err)
}
