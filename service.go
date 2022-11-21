package exswap

import (
	"math/big"
)

type AccountService interface {
    GetTokenBalanceOf(token string, userId int) *big.Int
}

type exchangeService struct {
}

func NewExchangeService() *exchangeService {
    s := new(exchangeService)
    return s
}

func (s *exchangeService) AddLiquidity() {

}

