package models

// 通证配置
type ExchangeToken struct {
    ConfigurableModel
    Name string `db:"name" json:"name"`
    Symbol string `db:"symbol" json:"symbol"`
    IconUrl string `db:"icon_url" json:"iconUrl"`
    Decimals int `db:"decimals" json:"decimals"`
    BizDecimals int `db:"biz_decimals" json:"bizDecimals"`
}
