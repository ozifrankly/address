package address

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidCEPFromCorreios(t *testing.T) {
	forever := make(chan Address)
	ctx, cancel := context.WithCancel(context.Background())
	go fetchCorreios(ctx, forever, "01001000")
	addr := <-forever
	cancel()
	assert.Equal(t, "São Paulo", addr.City)
	assert.Equal(t, "Sé", addr.District)
	assert.Equal(t, "- lado ímpar", addr.Complement)
	assert.Equal(t, "Praça da Sé", addr.Street)
	assert.Nil(t, addr.err)
}

func TestGetInvalidCEPFromCorreios(t *testing.T) {
	forever := make(chan Address)
	ctx, cancel := context.WithCancel(context.Background())
	go fetchCorreios(ctx, forever, "00000000")
	addr := <-forever
	cancel()
	assert.Error(t, addr.err)
}
