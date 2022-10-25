package main

import (
	"flag"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
    _ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/militch/exswap"
)


var (
	debug      bool = false
	host       string = "127.0.0.1"
	port       string = "9099"
	address    string
	config string = "./config.yml"
    db *sqlx.DB
    ko = koanf.New(".")
)

func init() {
    // 初始化命令行选项映射
	flag.StringVar(&address, "address", address, "Set listen address")
	flag.StringVar(&host, "host", host, "Set listen host")
	flag.StringVar(&port, "port", port, "Set listen port")
	flag.StringVar(&config, "config", config, "Set config file")
	flag.BoolVar(&debug, "debug", debug, "Set debug mode")
    // 初始化配置文件
    initConfig([]string{ config, }, ko)
    // 初始化数据库实例
    db = initDB()
}

func initConfig(files []string, ko *koanf.Koanf) {
    var err error
    var load bool
    for _, f := range files {
        logrus.Infof("Loading config file: %s", f)
        if err = ko.Load(file.Provider(f), yaml.Parser()); err != nil && !load {
            logrus.Errorf("Failed load config file: %v", err)
            panic(err)
        }else if err == nil && !load {
            load = true
        }
    }
}

func initDB() *sqlx.DB {
    var c struct {
        Host string `koanf:"host"`
        Port int `koanf:"port"`
        User string `koanf:"user"`
        Password string `koanf:"password"`
        Database string `koanf:"database"`
    }
    if err := ko.Unmarshal("mysql", &c); err != nil {
        logrus.Errorf("Failed parse mysql config: %v", err)
        panic(err)
    }
    dsn := fmt.Sprintf(
        "%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        c.User, c.Password, c.Host, c.Port, c.Database)
    logrus.Infof("Connecting to database: %s:%d/%s", 
        c.Host, c.Port, c.Database)
    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        logrus.Errorf("Failed connect to database: %v", err)
        panic(err)
    }
    return db
}

func main() {
	flag.Parse()
	srv := new(exswap.Server)
	addr := host + ":" + port
	if address != "" {
		addr = address
	}
	srv.ListenAddr = addr
	srv.Debug = debug
    srv.DB = db
    logrus.Infof("Start server listen on: %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
