package address

import (
	"context"
)

/*Fetch  makes a query in some services and returns the result of that ending first.
If all services cannot find a address, a error will be returned
*/
func Fetch(cep string) (*Address, error) {
	addrChan := make(chan Address)
	errorChan := make(chan error, 3)
	var addr Address
	count := 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go fetchCorreios(ctx, cep, addrChan, errorChan)
	go fetchWidenet(ctx, cep, addrChan, errorChan)
	go fetchViacep(ctx, cep, addrChan, errorChan)

	for {
		select {
		case addr = <-addrChan:
			return &addr, nil
		case err := <-errorChan:
			if count++; count == 3 {
				return nil, err
			}
		}
	}
}
