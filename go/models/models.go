package models

import "gorm.io/gorm"

// 客户基本信息表
type ClientsInfo struct {
    gorm.Model
    LimitBal  float64
    Sex       int
    Age       int
    Education int
    Marriage  int
    Defaulted int
}

// 性别分布表
type GenderInfo struct {
    gorm.Model
    Gender   int  `gorm:"unique" json:"gender"`
    Count int   `json:"count"`
}

// 年龄分布表
type AgeInfo struct {
    gorm.Model
    Age   int   `gorm:"unique" json:"age"`
    Count int   `json:"count"`
}

// 教育程度分布表
type EducationInfo struct {
    gorm.Model
    Education int   `gorm:"unique" json:"education"`
    Count     int   `json:"count"`
}

// 婚姻状况分布表
type MarriageInfo struct {
    gorm.Model
    Marriage int    `gorm:"unique" json:"marriage"`
    Count    int    `json:"count"`
}

// 信用额度范围分布表
type LimitBalInfo struct {
    gorm.Model
    Range string    `gorm:"unique" json:"range"`
    Count int       `json:"count"`
}

// 按性别违约统计表
type GenderDefaultsStats struct {
    gorm.Model
    Gender    int   `gorm:"uniqueIndex:idx_gender_default" json:"gender"`
    Defaulted int   `gorm:"uniqueIndex:idx_age_default" json:"defaulted"`
    Count     int   `json:"count"`
}

// 按年龄段违约统计表
type AgeDefaultsStats struct {
    gorm.Model
    Age       int   `gorm:"uniqueIndex:idx_age_default" json:"age"`
    Defaulted int   `gorm:"uniqueIndex:idx_age_default" json:"defaulted"`
    Count     int   `json:"count"`
}

// 按教育水平违约统计表
type EducationDefaultsStats struct {
    gorm.Model
    Education int   `gorm:"uniqueIndex:idx_education_default" json:"education"`
    Defaulted int   `gorm:"uniqueIndex:idx_education_default" json:"defaulted"`
    Count     int   `json:"count"`
}

// 按婚姻状况违约统计表
type MarriageDefaultsStats struct {
    gorm.Model
    Marriage  int   `gorm:"uniqueIndex:idx_marriage_default" json:"marriage"`
    Defaulted int   `gorm:"uniqueIndex:idx_marriage_default" json:"defaulted"`
    Count     int   `json:"count"`
}

// 按信用额度范围违约统计表
type LimitDefaultsStats struct {
    gorm.Model
    Range     string    `gorm:"varchar(255) uniqueIndex:idx_range_default" json:"range"`
    Defaulted int       `gorm:"uniqueIndex:idx_range_default" json:"defaulted"`
    Count     int       `json:"count"`
}

// 账单和还款对比表（按月份统计）
type BillPayComparison struct {
    gorm.Model
    Month     int       `gorm:"unique" json:"month"`
    TotalBill float64   `json:"totalnill"`
    TotalPay  float64   `json:"totalpay"`
}

// 连续违约统计表
type ConsecutiveDefaults struct {
    gorm.Model
    ClientID        uint    `gorm:"unique json:"client_id"`
    MaxConsecutive  int     `json:"maxconsecutive"`
}

// 综合违约评分汇总
type DefaultScoreSummary struct {
    gorm.Model
    ScoreName   string  `gorm:“unique" json:"scorename"`
    ClientCount int     `json:"clientcount"`
}

