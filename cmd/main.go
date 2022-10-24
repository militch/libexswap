package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/militch/exswap"
	"github.com/militch/exswap/global"
	"github.com/militch/exswap/initialize"
	"github.com/sirupsen/logrus"
)

const defaultConfigPath = "./config.yaml"

var (
	debug      bool
	host       string
	port       string
	address    string
	configPath string
)

func init() {
	flag.StringVar(&address, "address", "", "Set listen address")
	flag.StringVar(&host, "host", "127.0.0.1", "Set listen host")
	flag.StringVar(&port, "port", "9099", "Set listen port")
	flag.StringVar(&configPath, "config", defaultConfigPath, "")
	flag.BoolVar(&debug, "debug", false, "Set debug mode")
}

func getEnv() {
	klog.SetLogger(klog.NewStdLogger(io.Discard))
	c := config.New(
		config.WithSource(
			env.NewSource(),
			file.NewSource(configPath),
		))
	if err := c.Load(); err != nil {
		logrus.Errorf("c.Load %v", err)
	}

	if err := c.Scan(&global.GVA_CONFIG); err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	getEnv()
	global.GVA_DB = initialize.Gorm()
	srv := new(exswap.Server)
	addr := host + ":" + port
	if address != "" {
		addr = address
	}
	srv.ListenAddr = addr
	srv.Debug = debug
	fmt.Printf("Started server listen on: %s\n", addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
