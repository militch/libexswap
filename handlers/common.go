package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paginaction struct {
    Data interface{} `json:"data"`
    Page int `json:"page"`
    PerPage int `json:"perPage"`
    Total int `json:"total"`
}

type Search struct {
    PerPage int
}

func getPagination(ctx *gin.Context, 
    perPage int) *Paginaction {

    page, _ := strconv.Atoi(ctx.Query("page"))
    pp, _ := strconv.Atoi(ctx.Query("per_page"))
    if pp > 0 {
        perPage = pp
    }
    if page < 1 {
        page = 1
    }
    return &Paginaction {
        Page: page,
        PerPage: perPage,
    }
}

func getSearch(ctx *gin.Context, perPage int) Search {
    pp, _ := strconv.Atoi(ctx.Query("per_page"))
    if pp > 0 {
        perPage = pp
    }
    return Search {
        PerPage: perPage,
    }
}

func SendError(ctx *gin.Context, code int, err error) {
    if err == nil {
        panic("err is nil")
    }
    var parsedErr *RespError
    ok := errors.As(err, &parsedErr)
    if !ok {
        parsedErr = &RespError{
            Err: err,
            Code: code,
        }
    }
    _ = ctx.Error(parsedErr)
}
// 发送分页数据
func SendPage(c *gin.Context, p *Paginaction){
    c.JSON(200, p)
}
