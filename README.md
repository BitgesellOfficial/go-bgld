## bgld

<img src="Icon.png" style="height: 60px;"/>

[![Go Reference](https://pkg.go.dev/badge/github.com/bitgesellofficial/go-bgld.svg)](https://pkg.go.dev/github.com/bitgesellofficial/go-bgld)

A Go client library wrapping the BGLd JSON-RPC API for the Bitgesell blockchain network.


## Installation

```sh
go get github.com/bitgesellofficial/go-bgld
```


## Usage
----
```go
package main

import (
	"log"

	bgld "github.com/bitgesellofficial/go-bgld"
)

const (
	serverHost       = "localhost"
	serverPort       = 8454
	user             = "user"
	password         = "password"
	useSSL           = false
	walletPassphrase = "wallet-passphrase"
)

func main() {
	bc, err := bgld.New(serverHost, serverPort, user, password, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	err = bc.WalletPassphrase(walletPassphrase, 3600)
	log.Println(err)

	err = bc.BackupWallet("/tmp/wallet.dat")
	log.Println(err)

	privKey, err := bc.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
	log.Println(err, privKey)
}
```
	
More examples are available in `examples/example.go`.

## Documentation

Click on the button below to access the full documentation:

[Go Reference for github.com/bitgesellofficial/go-bgld](https://pkg.go.dev/github.com/bitgesellofficial/go-bgld)



## Unit tests

More than 100 unit tests are made.

To run tests:

```sh
go test ./...
```

	Running Suite: BGLd Suite	
	=============================
	Random Seed: 1401120770
	Will run 112 of 112 specs

	•••••••••••••••••••••••••••••••••••
	Ran 112 of 112 Specs in 0.001 seconds
	SUCCESS! -- 112 Passed | 0 Failed | 0 Pending | 0 Skipped PASS

	Ginkgo ran in 10.856335553s
	Test Suite Passed
 



Todo
-----
* GetBlockTemplate
* sendrawtransaction
* signrawtransaction
* submitblock

##### Note on SSL support 

Note on ssl support : bgld library doesn't verify the server's certificate chain. That means that it accepts any certificate presented by the server and any host name in that certificate. In this mode, TLS is susceptible to man-in-the-middle attacks.
