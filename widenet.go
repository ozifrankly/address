package address

import (
	"context"
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

func fetchWidenet(ctx context.Context, addr chan Address, cep string) {
	var url = fmt.Sprint("https://cep.widenet.host/busca-cep/api/cep/", cep, ".json")
	var inner widenetAddress

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		addr <- Address{err: err}
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		addr <- Address{err: err}
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&inner)

	if err != nil {
		addr <- Address{err: err}
		return
	}

	if !inner.OK {
		addr <- Address{err: errors.New("invalid cep")}
		return
	}
	addr <- Address{City: inner.City, District: inner.District, Complement: "", Street: inner.Address, err: nil}
}
