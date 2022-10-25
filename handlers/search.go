package handlers

import "github.com/gin-gonic/gin"

// 通证搜索
func SearchTokens(ctx *gin.Context) {
    s := getSearch(ctx, 10)
    _ = s
}

// 交易池搜索
func SearchPools(ctx *gin.Context) {

}
