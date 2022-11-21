package models

type BaseTokenPair struct {
}

type ExchangeTokenPair struct {
    ConfigurableModel
    Name string `db:"name" json:"name"`
    TokenMain string `db:"token_main" json:"tokenMain"`
    TokenSecond string `db:"token_second" json:"tokenSecond"`
}
