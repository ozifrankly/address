package address

import (
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
	Fail        bool   `json:"erro"`
}

func fetchViacep(cep string) (*Address, error) {
	var url = fmt.Sprint("https://viacep.com.br/ws/", cep, "/json/")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var inner viacepAddress
	err = json.NewDecoder(resp.Body).Decode(&inner)
	if err != nil {
		return nil, err
	}
	if inner.Fail {
		return nil, errors.New("invalid cep")
	}
	address := Address{City: inner.Localidade, District: inner.Bairro, Complement: inner.Complemento, Street: inner.Logradouro}
	return &address, nil
}
