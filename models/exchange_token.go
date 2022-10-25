package models

// 通证配置
type ExchangeToken struct {
    ConfigurableModel
    Name string `db:"name"`
    Symbol string `db:"symbol"`
    IconUrl string `db:"icon_url"`
    Decimals int `db:"decimals"`
    BizDecimals int `db:"biz_decimals"`
}
