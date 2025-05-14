package models
import (
    "gorm.io/gorm"
)

type Guest struct {
    gorm.Model
    ID          uint
    Sex         string
    Education   string
    Age         uint
    Marriage    string
    Limit_Bal   uint
    Default     uint
}

type Default-Statistic struct {
    gorm.Model
    ID          uint    `gorm:"unique; not null; primarykey" json:"id"`
    EventName   string  `gorm:"unique; not null" json:"eventname"`
    Defaults    uint    `gorm:"unique; not null" json:"defaults"`
    NonDefaults uint    `gorm:"unique; not null" json:“nondefaults"`
}

type DefaultBySex {
    gorm.Model
    Status      uint
    Defaults    uint
    NonDefaults uint
}

type DefaultByAge {
    gorm.Model
    Status      uint
    Defaults    uint
    NonDefaults uint
}

type DefaultByEdu {
    gorm.Model
    Status      uint
    Defaults    uint
    NonDefaults uint
}

type DefaultByMari {
    gorm.Model
    Status      uint
    Defaults    uint
    NonDefaults uint
}

type DefaultByLim {
    gorm.Model
    Status      uint
    Defaults    uint
    NonDefaults uint
}
