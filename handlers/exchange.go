package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


// 通证列表查询
func ExchangeGetTokens(ctx *gin.Context) {
    page := getPagination(ctx, 10)
    logrus.Infof("get tokens: %v", page)
    GetDB().Get()
    _ = page
    SendPage(ctx, page, nil)
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
