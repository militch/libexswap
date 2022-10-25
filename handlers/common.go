package handlers

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type Paginaction struct {
    Page int
    PerPage int
}

type Search struct {
    PerPage int
}

func getPagination(ctx *gin.Context, 
    perPage int) Paginaction {

    page, _ := strconv.Atoi(ctx.Query("page"))
    pp, _ := strconv.Atoi(ctx.Query("per_page"))
    if pp > 0 {
        perPage = pp
    }
    if page < 1 {
        page = 1
    }
    return Paginaction {
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

// 发送分页数据
func SendPage(c *gin.Context, p Paginaction, 
    data []interface{}){
    c.JSON(200, gin.H{
        "page": "abc",
    })
}
