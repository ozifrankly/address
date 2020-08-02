package address

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type viacepAddress struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Fail        bool   `json:"erro"`
}

func fetchViacep(ctx context.Context, cep string, addr chan Address, errChan chan error) {
	var inner viacepAddress
	var url = fmt.Sprint("https://viacep.com.br/ws/", cep, "/json/")

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		errChan <- err
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&inner)
	if err != nil {
		errChan <- err
		return
	}
	if inner.Fail {
		errChan <- errors.New("invalid cep")
		return
	}

	addr <- Address{City: inner.Localidade, District: inner.Bairro, Complement: inner.Complemento, Street: inner.Logradouro, State: inner.UF}
}
