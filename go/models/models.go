package models

import "gorm.io/gorm"

// 客户基本信息表
type ClientsInfo struct {
    gorm.Model
    ClientID  int     `gorm:"unique;column:client_id" json:"client_id"` // 显式主键

    LimitBal  float64 `json:"limit_bal"`   // 信用额度
    Sex       int     `json:"sex"`         // 性别（1=男，2=女）
    Education int     `json:"education"`   // 教育程度
    Marriage  int     `json:"marriage"`    // 婚姻状态
    Age       int     `json:"age"`         // 年龄

    Pay0 int `json:"pay_0"` // 还款状态
    Pay2 int `json:"pay_2"`
    Pay3 int `json:"pay_3"`
    Pay4 int `json:"pay_4"`
    Pay5 int `json:"pay_5"`
    Pay6 int `json:"pay_6"`

    BillAmt1 float64 `json:"bill_amt1"` // 账单金额
    BillAmt2 float64 `json:"bill_amt2"`
    BillAmt3 float64 `json:"bill_amt3"`
    BillAmt4 float64 `json:"bill_amt4"`
    BillAmt5 float64 `json:"bill_amt5"`
    BillAmt6 float64 `json:"bill_amt6"`

    PayAmt1 float64 `json:"pay_amt1"` // 还款金额
    PayAmt2 float64 `json:"pay_amt2"`
    PayAmt3 float64 `json:"pay_amt3"`
    PayAmt4 float64 `json:"pay_amt4"`
    PayAmt5 float64 `json:"pay_amt5"`
    PayAmt6 float64 `json:"pay_amt6"`

    Defaulted int `json:"defaulted"` // 是否违约（1=是，0=否）
}

type User struct {
    gorm.Model
    Username string `gorm:"unique; not null" json:username"`
    Password string `gorm:"not null" json:"-"`
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
    Defaulted int   `gorm:"uniqueIndex:idx_gender_default" json:"defaulted"`
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
    Range     string    `gorm:"type:varchar(255);uniqueIndex:idx_range_default" json:"range"`
    Defaulted int       `gorm:"uniqueIndex:idx_range_default" json:"defaulted"`
    Count     int       `json:"count"`
}

// 账单和还款对比表（按月份统计）
type BillPayComparison struct {
    gorm.Model
    Month     int       `gorm:"unique" json:"month"`
    TotalBill float64   `json:"totalbill"`
    TotalPay  float64   `json:"totalpay"`
}

// 连续违约统计表
type ConsecutiveDefaults struct {
    gorm.Model
    ClientID        uint    `gorm:"unique" json:"client_id"`
    MaxConsecutive  int     `json:"maxconsecutive"`
}

// 综合违约评分汇总
type DefaultScoreSummary struct {
    gorm.Model
    ScoreName   string  `gorm:"unique" json:"scorename"`
    ClientCount int     `json:"clientcount"`
}

type BillAndRepayment struct {
    Month string    `gorm:"unique" json:"month"`
    AverageBill float64 `json:"averageBill"`
    AverageRepay   float64 `json:"averageRepay"`
    OverdueAvgBill float64 `json:"overdueAvgBill"`
    NormalRepayRate float64 `json:"normalRepayRate"`
}

type OverdueStreakInfo struct {
    MaxOverdueStreak int    `json:"max_overdue_streak"`
    CustomerCount    int    `json:"customer_count"`
}

