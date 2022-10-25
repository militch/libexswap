package models

import "time"

type BaseModel struct {
    ID int `db:"id" json:"id"`
    CreateTime time.Time `db:"create_time" json:"createTime"`
    UpdateTime time.Time `db:"update_time" json:"updateTime"`
}

type WeightModel struct {
    Weight int `db:"weight" json:"weight"`
}

type DescriptionModel struct {
    Description string `db:"description" json:"description"`
}

type DisabledModel struct {
    IsDisabled bool `db:"is_disabled" json:"isDisabled"`
}

type ConfigurableModel struct {
    BaseModel
    WeightModel
    DisabledModel
    DescriptionModel
}
