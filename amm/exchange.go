package amm

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/militch/libexswap/common"
)

type Exchange interface {
    AddLiquidity(*big.Int, *big.Int, *big.Int) (*big.Int,error);
}

type exchange struct {
    totalSupply *big.Int
    balance *big.Int
    tokenBalance *big.Int
}

func NewExchange() *exchange {
    return &exchange{}
}

func (e *exchange) SetTotalSupply(num *big.Int) {
    e.totalSupply = num
}

func (e *exchange) SetBalance(num *big.Int) {
    e.balance = num
}

func (e *exchange) SetTokenBalance(num *big.Int) {
    e.tokenBalance = num
}

func (e *exchange) addTokenBalance(num *big.Int) {
    e.tokenBalance = new(big.Int).Add(e.tokenBalance, num)
}

func (e *exchange) subTokenBalance(num *big.Int) {
    e.tokenBalance = new(big.Int).Sub(e.tokenBalance, num)
}

func (e *exchange) AddLiquidity(value *big.Int,
    minLiquidity *big.Int, maxTokens *big.Int) (*big.Int, error) {
    totalLiquidity := e.totalSupply
    // totalLiquidity > 0
    if common.BigIntCmpGt(totalLiquidity, common.BIG_ZERO) {
        // assert minLiquidity > 0
        if (common.BigIntCmpLte(minLiquidity, common.BIG_ZERO)) {
            return nil, errors.New("minLiquidity must be > 0") 
        }
        deltaReserve := new(big.Int).Sub(e.balance, value)
        tokenReserve := e.tokenBalance
        // tokenAmount = value * tokenReserve / deltaReserve + 1
        tokenAmount := new(big.Int).Mul(value, tokenReserve)
        tokenAmount = new(big.Int).Div(tokenAmount, deltaReserve)
        tokenAmount = new(big.Int).Add(tokenAmount, common.BIG_ONE)
        // liquidityMinted = value * totalLiquidity / deltaReserve
        liquidityMinted := new(big.Int).Mul(value, totalLiquidity)
        liquidityMinted = new(big.Int).Div(liquidityMinted, deltaReserve)
        // assert maxTokens >= tokenAmount && liquidityMinted >= minLiquidity
        if common.BigIntCmpLt(maxTokens, tokenAmount) || 
            common.BigIntCmpLt(liquidityMinted, minLiquidity) {
            return liquidityMinted, fmt.Errorf("failed add liquidity")
        }
        e.addTokenBalance(tokenAmount)
        e.totalSupply = new(big.Int).Add(totalLiquidity, liquidityMinted)
        return liquidityMinted, nil
    }
    tokenAmount := maxTokens
    initialLiquidity := e.balance
    e.addTokenBalance(tokenAmount)
    e.totalSupply = initialLiquidity
    return initialLiquidity, nil
}

func (e *exchange) GetInputPrice(inputAmount *big.Int, inputReserve *big.Int,
    outputReserve *big.Int) (*big.Int, error) {
    // assert inputReserve > 0 && outputReserve > 0
    if common.BigIntCmpLte(inputReserve, common.BIG_ZERO) ||
        common.BigIntCmpLte(outputReserve, common.BIG_ZERO) {
        return nil, errors.New("inputReserve and outputReserve must be > 0")
    }
    inputAmountWithFee := new(big.Int).Mul(inputAmount, big.NewInt(997))
    numerator := new(big.Int).Mul(inputAmountWithFee, outputReserve)
    denominator := new(big.Int).Mul(inputReserve, big.NewInt(1000))
    denominator = new(big.Int).Add(denominator, inputAmountWithFee)
    return new(big.Int).Div(numerator, denominator), nil
}

func (e *exchange) GetOutputPrice(outputAmount *big.Int, 
    inputReserve *big.Int, outputReserve *big.Int) (*big.Int, error) {
    // assert inputReserve > 0 && outputReserve > 0
    if common.BigIntCmpLte(inputReserve, common.BIG_ZERO) ||
        common.BigIntCmpLte(outputReserve, common.BIG_ZERO) {
        return nil, errors.New("inputReserve and outputReserve must be > 0")
    }
    numerator := new(big.Int).Mul(inputReserve, outputAmount)
    numerator = new(big.Int).Mul(numerator, big.NewInt(1000))
    denominator := new(big.Int).Sub(outputReserve, outputAmount)
    denominator = new(big.Int).Mul(denominator, big.NewInt(997))
    return new(big.Int).Div(numerator, denominator), nil
}

func (e *exchange) DeltaToTokenInput(deltaSold *big.Int, minTokens *big.Int) (*big.Int, error) {
    // assert deltaSold > 0 && minTokens > 0
    if common.BigIntCmpLte(deltaSold, common.BIG_ZERO) ||
        common.BigIntCmpLte(minTokens, common.BIG_ZERO) {
        return nil, errors.New("deltaSold and minTokens must be > 0")
    }
    tokenReserve := e.tokenBalance
    deltaReserve := new(big.Int).Sub(e.balance, deltaSold)
    tokensBought, err := e.GetInputPrice(deltaSold, deltaReserve, tokenReserve)
    if err != nil {
        return nil, err
    }
    // assert tokensBought >= minTokens
    if common.BigIntCmpLt(tokensBought, minTokens) {
        return nil, fmt.Errorf("tokenBought(%s) must be >= minTokens(%s)", tokensBought, minTokens)
    }
    e.subTokenBalance(tokensBought)
    return tokensBought, nil
}

func (e *exchange) DeltaToTokenOutput(tokensBought *big.Int, maxDelta *big.Int) (*big.Int, error) {
    // assert tokensBought > 0 && maxDelta > 0
    if common.BigIntCmpLte(tokensBought, common.BIG_ZERO) ||
        common.BigIntCmpLte(maxDelta, common.BIG_ZERO) {
        return nil, errors.New("tokensBought and maxDelta must be > 0")
    }
    tokenReserve := e.tokenBalance
    deltaReserve := new(big.Int).Sub(e.balance, maxDelta)
    tokensBought, err := e.GetOutputPrice(tokensBought, deltaReserve, tokenReserve)
    if err != nil {
        return nil, err
    }
    // assert tokensBought >= minTokens
    if common.BigIntCmpLt(tokensBought, minTokens) {
        return nil, fmt.Errorf("tokenBought(%s) must be >= minTokens(%s)", tokensBought, minTokens)
    }
    e.subTokenBalance(tokensBought)
    return tokensBought, nil
}
