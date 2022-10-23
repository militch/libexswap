package main

import (
	"fmt"
	"github.com/militch/libexswap/amm"
)

func main(){ 
    ex := amm.NewExchange()
    _ = ex
    //ex.AddLiquidity()
    fmt.Printf("hello, word\n")
}
