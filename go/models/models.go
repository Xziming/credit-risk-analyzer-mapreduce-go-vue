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
type SexInfo struct {
    gorm.Model
    Sex   string
    Count int
}

// 年龄分布表
type AgeInfo struct {
    gorm.Model
    Age   int
    Count int
}

// 教育程度分布表
type EducationInfo struct {
    gorm.Model
    Education string
    Count     int
}

// 婚姻状况分布表
type MarriageInfo struct {
    gorm.Model
    Marriage string
    Count    int
}

// 信用额度范围分布表
type LimitBalRangeInfo struct {
    gorm.Model
    Range string
    Count int
}

// 按性别违约统计表
type GenderDefaultsStats struct {
    gorm.Model
    Sex       string
    Defaulted int // 0或1
    Count     int
}

// 按年龄段违约统计表
type AgeDefaultsStats struct {
    gorm.Model
    Age       int
    Defaulted int
    Count     int
}

// 按教育水平违约统计表
type EducationDefaultsStats struct {
    gorm.Model
    Education string
    Defaulted int
    Count     int
}

// 按婚姻状况违约统计表
type MarriageDefaultsStats struct {
    gorm.Model
    Marriage  string
    Defaulted int
    Count     int
}

// 按信用额度范围违约统计表
type LimitRangeDefaultsStats struct {
    gorm.Model
    Range     string
    Defaulted int
    Count     int
}

// 账单和还款对比表（按月份统计）
type BillPayComparison struct {
    gorm.Model
    Month     int
    TotalBill float64
    TotalPay  float64
}

// 连续违约统计表
type ConsecutiveDefaults struct {
    gorm.Model
    ClientID        uint
    MaxConsecutive  int
}

// 综合违约评分汇总
type DefaultScoreSummary struct {
    gorm.Model
    ScoreName   string // 自定义评分项（例如："高风险客户", "低风险" 等）
    ClientCount int
}

