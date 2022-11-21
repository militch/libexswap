package handlers

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/militch/exswap/common"
	"github.com/militch/exswap/models"
	"github.com/sirupsen/logrus"
)

type ExchangePoolObj struct {
    ID int `json:"id"`
    Name string `json:"name"`
    TokenMain string `json:"tokenMain"`
    TokenSecond string `json:"tokenSecond"`
    MainBalance *string `json:"mainBalance"`
    SecondBalance *string `json:"secondBalance"`
    TotalSupply *string `json:"totalSupply"`
    Weight int `json:"weight"`
    Description string `json:"description"`
    IsDisabled bool `json:"isDisabled"`
    CreateTime time.Time `json:"createTime"`
    UpdateTime time.Time `json:"updateTime"`
}

func GetTokenPair(name string) (*models.ExchangeTokenPair, error) {
    var tokenPair models.ExchangeTokenPair
    err := GetOne("exchange_token_pair", &tokenPair, 
        "WHERE name = ?", name)
    if err != nil && err == sql.ErrNoRows{
        return nil, nil
    } 
    return &tokenPair, err
}

func getPoolByName(name string) (*models.ExchangePool, error) {
    var pool models.ExchangePool
    err := GetOne("exchange_pool", &pool, 
        "WHERE name = ?", name)
    if err != nil && err == sql.ErrNoRows{
        return nil, nil
    } 
    return &pool, err
}

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
    page := getPagination(ctx, 10)
    var err error
    var tokenPairs []models.ExchangeTokenPair
    page, err = SelectPage(
        "exchange_token_pair",
        page, &tokenPairs, 
        "ORDER BY weight DESC")
    if err != nil {
        SendError(ctx, 0, err)
        return
    }
    var objs []ExchangePoolObj = make([]ExchangePoolObj, 0)
    for _, item := range tokenPairs {
        pool, _ := getPoolByName(item.Name)
        var obj ExchangePoolObj
        common.SimpleObjCopy(pool, &obj)
        common.SimpleObjCopy(item, &obj)
        objs = append(objs, obj)
    }
    page.Data = objs
    SendPage(ctx, page)
}

func ExchangeGetPoolDetail(ctx *gin.Context){

    //getTokenPair()
}


func ExchangeAddPoolLiquidity(ctx *gin.Context) {
    logrus.Infof("-- addliquidity")
    name := ctx.PostForm("name")
    logrus.Infof("name: %s", name)
    pair, err := GetTokenPair(name)
    if err != nil {
        SendError(ctx, 0, err)
        return
    }
    logrus.Infof("pair: %v", pair)
    pool, err := getPoolByName(name)
    if err != nil {
        SendError(ctx, 0, err)
        return
    }
    _ = pool

    logrus.Infof("addliquidity: %v", pool)
}

func ExchangeRemovePoolLiquidity(ctx *gin.Context){

    name := ctx.PostForm("name")
    pool, err := getPoolByName(name)
    if err != nil {
        SendError(ctx, 0, err)
        return
    }
    _ = pool

    logrus.Infof("remove: %v", pool)
}

func ExchangeTrade(ctx *gin.Context) {

}

func ExchangeTestTrade(ctx *gin.Context) {
    //ex := 
}
