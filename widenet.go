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

func fetchWidenet(ctx context.Context, cep string, addr chan Address, errChan chan error) {
	var url = fmt.Sprint("https://cep.widenet.host/busca-cep/api/cep/", cep, ".json")
	var inner widenetAddress

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

	if !inner.OK {
		errChan <- errors.New("invalid cep")
		return
	}
	addr <- Address{City: inner.City, District: inner.District, Complement: "", Street: inner.Address, State: inner.State}
}
