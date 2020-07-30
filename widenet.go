package address

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type widenetAddress struct {
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
	OK       bool   `json:"ok"`
}

func fetchWidenet(cep string) (*Address, error) {
	var url = fmt.Sprint("https://cep.widenet.host/busca-cep/api/cep/", cep, ".json")
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
	var inner widenetAddress
	err = json.NewDecoder(resp.Body).Decode(&inner)
	if err != nil {
		return nil, err
	}
	if !inner.OK {
		return nil, errors.New("invalid cep")
	}
	address := Address{City: inner.City, District: inner.District, Complement: "", Street: inner.Address}
	return &address, nil
}
