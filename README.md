Address
================================

ℹ️ A search for the CEP integrates with the Correios, viaCEP and Widenet.

[![Build Status](https://travis-ci.org/ozifrankly/address.svg)](https://travis-ci.org/ozifrankly/address) [![Go Report Card](https://goreportcard.com/badge/github.com/ozifrankly/address)](https://goreportcard.com/report/github.com/ozifrankly/address) [![GoDoc](https://godoc.org/github.com/ozifrankly/address?status.svg)](https://godoc.org/github.com/ozifrankly/address)


## Features
  - Always up to date because it connects to three services directly.
  - Has high availability for using various services.
  - It is quick to make queries concurrently.
  
## Usage


```go
package main

import (
	"fmt"

	"github.com/ozifrankly/address"
)

func main() {
	addr, err := address.Fetch("01001000")
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
}

```

## Response

```
Address {
	State      "SP"
	City       "São Paulo"
	District   "Sé"
	Street     "Praça da Sé"
	Complement "lado ímpar"
}
```
