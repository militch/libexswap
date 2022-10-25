package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/militch/exswap/models"
)

// 通证列表查询
func ExchangeGetTokens(ctx *gin.Context) {
    page := getPagination(ctx, 10)
    var err error
    var tokens []models.ExchangeToken
    page, err = SelectPage(
        "exchange_token",
        page, &tokens, 
        "ORDER BY weight DESC")
    if err != nil {
        SendError(ctx, 0, err)
        return
    }
    SendPage(ctx, page)
}

func ExchangeGetPools(ctx *gin.Context) {

}

func ExchangeGetPoolDetail(ctx *gin.Context){

}


func ExchangeAddPoolLiquidity(ctx *gin.Context) {

}

func ExchangeRemovePoolLiquidity(ctx *gin.Context){

}

func ExchangeTrade(ctx *gin.Context) {

}

func ExchangeTestTrade(ctx *gin.Context) {

}
