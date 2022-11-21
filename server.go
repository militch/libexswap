package exswap

import (
	"net"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/militch/exswap/core"
	"github.com/militch/exswap/handlers"
)

type Server struct {
    ginapp  *gin.Engine
    ListenAddr string
    Debug bool
    DB *sqlx.DB
}

func (srv *Server) initialRouters(){
    ginapp := srv.ginapp
    srvCtx := &core.ServerContext{
        DB: srv.DB,
    }
    handlers.SrvContext = srvCtx
    // AMM 开发模块

    developGroup := ginapp.Group(
        "/develop")
    // 流动性趋势
    developGroup.GET("/liquidity", 
        handlers.AmmLiquidity)

    // 账户相关模块
    accountGroup := ginapp.Group(
        "/account")
    accountGroup.GET("/assets", 
        handlers.AccountGetAssets)
    accountGroup.GET("/logs", 
        handlers.AccountGetLogs)
    accountGroup.POST("/transfer", 
        handlers.AccountTransfer)

    // 市场相关模块
    marketGroup := ginapp.Group(
        "/market")
    marketGroup.GET("/latest", 
        handlers.MarketGetLatest)
    marketGroup.GET("/history", 
        handlers.MarketGetHistory)
    marketGroup.GET("/trades", 
        handlers.MarketGetTrades)

    // 交易相关模块
    exchangeGroup := ginapp.Group(
        "/exchange")
    exchangeGroup.GET("/tokens", 
        handlers.ExchangeGetTokens)
    exchangeGroup.GET("/pools", 
        handlers.ExchangeGetPools)
    exchangeGroup.GET("/pools/:name", 
        handlers.ExchangeGetPoolDetail)
    exchangeGroup.POST("/liquidity", 
        handlers.ExchangeAddPoolLiquidity)
    exchangeGroup.DELETE("/liquidity", 
        handlers.ExchangeRemovePoolLiquidity)
    exchangeGroup.POST("/trade", 
        handlers.ExchangeTrade)
    exchangeGroup.POST("/trade/test", 
        handlers.ExchangeTestTrade)

    // 业务搜索模块
    searchGroup := ginapp.Group(
        "/search")
    searchGroup.GET("/tokens", 
        handlers.SearchTokens)
    searchGroup.GET("/pools", 
        handlers.SearchPools)
}

func (srv *Server) ListenAndServe() error {
    if !srv.Debug {
        gin.SetMode(gin.ReleaseMode)
    }
    ginapp := gin.Default()
    srv.ginapp = ginapp
    ginapp.NoRoute(handlers.HandleNoRoute())
    ginapp.Use(cors.Default())
    ginapp.Use(handlers.HandleErrors())
    srv.initialRouters()
    ln, err := net.Listen("tcp", srv.ListenAddr)
	if err != nil {
		return err
	}
    return ginapp.RunListener(ln)
}
