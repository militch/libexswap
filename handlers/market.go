package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func MarketGetLatest(ctx *gin.Context) {
    // m, _ := strconv.Atoi("0")
    // n := 1 / m
    // _ = n
    SendRespError(ctx, 100, errors.New("dd"))
}
func MarketGetHistory(ctx *gin.Context) {
}

func MarketGetTrades(ctx *gin.Context) {
}
