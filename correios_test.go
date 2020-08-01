package address

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEPFromCorreios(t *testing.T) {
	var err error
	var addr Address
	forever := make(chan Address)
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go fetchWidenet(ctx, "01001000", forever, errChan)

	select {
	case addr = <-forever:
	case err = <-errChan:
	}

	assert.Equal(t, "São Paulo", addr.City)
	assert.Equal(t, "Sé", addr.District)
	assert.Equal(t, "", addr.Complement)
	assert.Equal(t, "Praça da Sé - lado ímpar", addr.Street)
	assert.Nil(t, err)
}

func TestGetInvalidCEPFromCorreios(t *testing.T) {
	var err error
	var addr Address
	forever := make(chan Address)
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go fetchCorreios(ctx, "00000000", forever, errChan)

	select {
	case addr = <-forever:
	case err = <-errChan:
	}

	assert.Empty(t, addr)
	assert.NotNil(t, err)
	assert.Error(t, err)
}
