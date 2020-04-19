# go-btcturk

golang client for btcturk api

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/aliereno/go-btcturk)](https://goreportcard.com/report/github.com/aliereno/go-btcturk)
[![HitCount](http://hits.dwyl.com/aliereno/go-btcturk.svg)](http://hits.dwyl.com/aliereno/go-btcturk)

</div>

#### This repo has been forked from [vural/go-btcturk](https://github.com/vural/go-btcturk) and has been improved for Btcturk's new api. So if you like this repo please give a star actual [repo](https://github.com/vural/go-btcturk).

## documentation

the documentation is available on [godoc](http://godoc.org/github.com/aliereno/go-btcturk/btcturk)

## install

```sh
go get -u github.com/aliereno/go-btcturk/btcturk
```

## usage
```go
package main

import (
	"github.com/aliereno/go-btcturk/btcturk"
)

func main() {
    api := btcturk.NewBTCTurkClient()
    t, err := api.Ticker()
    if err != nil {
        print(err)
        os.Exit(1)
    }
    
    for _, v := range t {
        println(v.Ask)
    }

    // if you don't plan to call authenticated api methods. SetAuthKey not required.
    api.SetAuthKey("publicKey", "privateKey")
}

```

## Examples

```go
package main

import (
	"github.com/aliereno/btcturk/btcturk"
)

func main() {
    api := btcturk.NewBTCTurkClient()

    //PUBLIC ENDPOINTS

    //TICKER
    _, _ = api.PairSymbol(btcturk.BTCTRY).Ticker()

    //ORDER BOOK
    _, _ = api.PairSymbol(btcturk.BTCTRY).Limit(10).OrderBook() // limit optional

    //TRADES
    _, _ = api.PairSymbol(btcturk.BTCTRY).Trades()


    //PRIVATE ENDPOINTS

    _, _ = api.SetAuthKey("publicKey", "privateKey")

    _, _ = api.Balance()

    _, _ = api.Quantity(0.001).
        Price(50000).
        StopPrice(0).
        OrderMethod("limit").
        PairSymbol(btcturk.BTCTRY).
        Buy()
}

```

## Notes
 - you can get your private/public key peer from your account

**[BTCTurk API documentation](https://docs.btcturk.com)**
