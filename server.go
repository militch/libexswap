package exswap

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/militch/exswap/handles"
)

type Server struct {
    ginapp  *gin.Engine
    ListenAddr string
    Debug bool
}

func (srv *Server) initialRouters(){
    ginapp := srv.ginapp
    accountGroup := ginapp.Group("/accounts")
    accountGroup.GET("/assets", 
        handles.AccountGetAssets)
    accountGroup.GET("/logs", 
        handles.AccountGetLogs)
    accountGroup.POST("/transfer", 
        handles.AccountTransfer)

    marketGroup := ginapp.Group("/market")
    marketGroup.GET("/latest", 
        handles.MarketGetLatest)
    marketGroup.GET("/history", 
        handles.MarketGetLatest)
    marketGroup.GET("/trades", 
        handles.MarketGetTrades)

    exchangeGroup := ginapp.Group("/exchange")
    exchangeGroup.GET("/tokens", 
        handles.ExchangeGetTokens)
    exchangeGroup.GET("/pools", 
        handles.ExchangeGetPools)
    exchangeGroup.GET("/pools/:name", 
        handles.ExchangeGetPoolDetail)
    exchangeGroup.POST("/liquidity", 
        handles.ExchangeAddPoolLiquidity)
    exchangeGroup.DELETE("/liquidity", 
        handles.ExchangeRemovePoolLiquidity)
    exchangeGroup.POST("/trade", 
        handles.ExchangeTrade)
    exchangeGroup.POST("/trade/test", 
        handles.ExchangeTestTrade)
}

func (srv *Server) ListenAndServe() error {
    if !srv.Debug {
        gin.SetMode(gin.ReleaseMode)
    }
    ginapp := gin.Default()
    srv.ginapp = ginapp
    srv.initialRouters()
    ln, err := net.Listen("tcp", srv.ListenAddr)
	if err != nil {
		return err
	}
    return ginapp.RunListener(ln)
}
