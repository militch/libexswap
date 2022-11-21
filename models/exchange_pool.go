package models

import "math/big"

type ExchangePool struct {
    BaseModel
    Name string `db:"name" json:"name"`
    MainBalance *big.Int `db:"main_balance" json:"mainBalance"`
    SecondBalance *big.Int `db:"second_balance" json:"secondBalance"`
    TotalSupply *big.Int `db:"total_supply" json:"totalSupply"`
}
