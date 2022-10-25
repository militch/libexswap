package models

type ExchangeTokenPair struct {
    ConfigurableModel
    Name string `db:"name"`
    TokenMain string `db:"token_main"`
    TokenSecond string `db:"token_second"`
}
