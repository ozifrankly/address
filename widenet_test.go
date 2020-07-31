package address

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEPFromWidenet(t *testing.T) {
	forever := make(chan Address)
	ctx, cancel := context.WithCancel(context.Background())
	go fetchWidenet(ctx, forever, "01001000")
	address := <-forever
	cancel()
	assert.Equal(t, "São Paulo", address.City)
	assert.Equal(t, "Sé", address.District)
	assert.Equal(t, "", address.Complement)
	assert.Equal(t, "Praça da Sé - lado ímpar", address.Street)
	assert.Nil(t, address.err)
}

func TestGetInvalidCEPFromWidenet(t *testing.T) {
	forever := make(chan Address)
	ctx, cancel := context.WithCancel(context.Background())
	go fetchWidenet(ctx, forever, "00000000")
	address := <-forever
	cancel()
	assert.Error(t, address.err)
}
