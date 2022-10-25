package handlers

import (
    "github.com/jmoiron/sqlx"
	"github.com/militch/exswap/core"
)

var (
    SrvContext *core.ServerContext
)

func GetDB() *sqlx.DB {
    return SrvContext.DB
}
