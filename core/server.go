package core

import "github.com/jmoiron/sqlx"

type ServerContext struct {
    DB *sqlx.DB
}
