package models

import "time"

type BaseModel struct {
    ID int `db:"id"`
    CreateTime time.Time `db:"create_time"`
    UpdateTime time.Time `db:"update_time"`
}

type WeightModel struct {
    Weight int `db:"weight"`
}

type DescriptionModel struct {
    Description string `db:"description"`
}

type DisabledModel struct {
    IsDisabled bool `db:"is_disabled"`
}

type ConfigurableModel struct {
    BaseModel
    WeightModel
    DisabledModel
    DescriptionModel
}
