package handlers

import "github.com/gin-gonic/gin"

func ExchangeGetTokens(ctx *gin.Context) {
    page := getPagination(ctx, 10)
    _ = page
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
