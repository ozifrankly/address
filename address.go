package address

import (
	"context"
)

func fetch(cep string) *Address {
	forever := make(chan Address)
	var addr Address
	ctx, cancel := context.WithCancel(context.Background())
	go fetchCorreios(ctx, forever, cep)
	go fetchWidenet(ctx, forever, cep)
	go fetchViacep(ctx, forever, cep)
	addr = <-forever
	cancel()
	return &addr
}
