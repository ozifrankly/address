package address

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *soapBodyResponse
}
type soapBodyResponse struct {
	XMLName xml.Name `xml:"Body"`
	Resp    *responseBody
}

type responseBody struct {
	XMLName  xml.Name `xml:"http://cliente.bean.master.sigep.bsb.correios.com.br/ consultaCEPResponse"`
	Response *body
}

type body struct {
	XMLName    xml.Name `xml:"return"`
	District   string   `xml:"bairro"`
	Code       string   `xml:"cep"`
	City       string   `xml:"cidade"`
	Complement string   `xml:"complemento2"`
	Address    string   `xml:"end"`
	State      string   `xml:"uf"`
}

const xmlBegin = "<?xml version=\"1.0\"?><soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:cli=\"http://cliente.bean.master.sigep.bsb.correios.com.br/\"><soapenv:Header /><soapenv:Body><cli:consultaCEP><cep>"
const xmlEnd = "</cep></cli:consultaCEP></soapenv:Body></soapenv:Envelope>"

func fetchCorreios(ctx context.Context, addr chan Address, cep string) {
	var inner response
	var bodyBytes []byte
	var url = "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente"
	var data = []byte(fmt.Sprint(xmlBegin, cep, xmlEnd))

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		addr <- Address{err: err}
	}

	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		addr <- Address{err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 500 {
		addr <- Address{err: errors.New("invalid cep")}
		return
	}

	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		addr <- Address{err: err}
		return
	}

	if err := xml.Unmarshal(bodyBytes, &inner); err != nil {
		addr <- Address{err: err}
		return
	}
	resposeAddr := inner.SoapBody.Resp.Response

	addr <- Address{City: resposeAddr.City, District: resposeAddr.District, Complement: resposeAddr.Complement, Street: resposeAddr.Address}
}
