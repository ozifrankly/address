package address

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEPFromViacep(t *testing.T) {
	var err error
	var addr Address
	forever := make(chan Address)
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go fetchViacep(ctx, "01001000", forever, errChan)

	select {
	case addr = <-forever:
	case err = <-errChan:
	}

	assert.Equal(t, "SP", addr.State)
	assert.Equal(t, "São Paulo", addr.City)
	assert.Equal(t, "Sé", addr.District)
	assert.Equal(t, "lado ímpar", addr.Complement)
	assert.Equal(t, "Praça da Sé", addr.Street)
	assert.Nil(t, err)
}

func TestGetInvalidCEPFromViacep(t *testing.T) {
	var err error
	var addr Address
	forever := make(chan Address)
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go fetchViacep(ctx, "00000000", forever, errChan)

	select {
	case addr = <-forever:
	case err = <-errChan:
	}

	assert.Empty(t, addr)
	assert.NotNil(t, err)
	assert.Error(t, err)
}
