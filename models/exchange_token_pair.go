package models

type ExchangeTokenPair struct {
    ConfigurableModel
    Name string `db:"name" json:"name"`
    TokenMain string `db:"token_main" json:"token_main"`
    TokenSecond string `db:"token_second" json:"token_second"`
}
