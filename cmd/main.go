package main

import (
	"flag"
	"fmt"

	"github.com/militch/exswap"
)

var (
    debug bool
    host string
    port string
    address string
)

func init(){
    flag.StringVar(&address, "address", "", "Set listen address")
    flag.StringVar(&host, "host", "127.0.0.1", "Set listen host")
    flag.StringVar(&port, "port", "9099", "Set listen port")
    flag.BoolVar(&debug, "debug", false, "Set debug mode")
}

func main(){ 
    flag.Parse()
    srv := new(exswap.Server)
    addr := host + ":" + port
    if address != "" {
        addr = address
    }
    srv.ListenAddr = addr
    srv.Debug = debug
    fmt.Printf("Started server listen on: %s\n", addr)
    if err := srv.ListenAndServe(); err != nil{
        panic(err)
    }
}
