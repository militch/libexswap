package handlers

import (
	"fmt"
)

func SelectPage(db string, p *Paginaction,
    dest interface{}, sql string, 
    args ...interface{}) (*Paginaction, error) {
    offset := p.Page - 1
    if offset < 0 {
        offset = 0
    }
    var err error
    var count int
    countsql := fmt.Sprintf("SELECT count(*) FROM `%s` %s;", db, sql)
    err = GetDB().Get(&count, countsql, args...)
    if err != nil {
        return p, err
    }
    p.Total = count
    sql = fmt.Sprintf("SELECT * FROM `%s` %s LIMIT %d OFFSET %d;", 
        db, sql, p.PerPage, offset)
    err = GetDB().Select(dest, sql, args...)
    if err != nil {
        return p, err
    }
    p.Data = dest
    return p, nil
}
