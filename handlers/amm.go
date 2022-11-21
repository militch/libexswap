package handlers

import (
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/militch/exswap/amm"
	"github.com/militch/exswap/common"
	"github.com/sirupsen/logrus"
)

type Liquidity struct {
    Value string `json:"value"`
    Liquidity string `json:"liquidity"`
    TokenAmount string `json:"tokenAmount"`
}

func yuantofen(amount *big.Int) *big.Int {
    return new(big.Int).Mul(amount, common.BIG_TEN_POW_EIGHT)
}

func fentoyuan(amount *big.Int) *big.Float {
    amountf := new(big.Float).SetInt(amount)
    per := new(big.Float).SetInt(common.BIG_TEN_POW_EIGHT)
    return new(big.Float).Quo(amountf, per)
}

var (
    INITIAL_DELTA = new(big.Int).SetUint64(uint64(100))
    INITIAL_TOKEN = new(big.Int).SetUint64(uint64(24500))
)

func AmmLiquidity(ctx *gin.Context){
    ex := amm.NewExchange()
    _ = ex
    var liquiditys []Liquidity
    initialValue := yuantofen(INITIAL_DELTA)
    // 初始预储值
    ex.AddBalance(initialValue)
    // 初始流动性
    initialLiquidity, _, _ := ex.AddLiquidity(
        initialValue, common.BIG_ONE, yuantofen(INITIAL_TOKEN))
    logrus.Infof("Initial liquidity: %s", initialLiquidity)
    for i:=uint64(200); i<=uint64(400); i++{
        value := new(big.Int).SetUint64(i)
        value = yuantofen(value)
        // 预储值
        ex.AddBalance(value)
        // 添加流动性
        liquidity, tokenAmount, err := ex.AddLiquidity(
            value, common.BIG_ONE, common.BIG_ONE)
        if err != nil {
            logrus.Warnf("failed add liquidity: %v", err)
            break
        }
        la := fentoyuan(liquidity)
        ta := fentoyuan(tokenAmount)
        liquiditys = append(liquiditys, Liquidity{
            Value: value.String(),
            Liquidity: la.String(),
            TokenAmount: ta.String(),
        })
        //ex.AddLiquidity()
    }
    SendData(ctx, liquiditys)
}
