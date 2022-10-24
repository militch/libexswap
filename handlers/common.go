package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)


type Paginaction struct {
    Page int
    PerPage int
}

func getPagination(ctx *gin.Context, perPage int) Paginaction {
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
