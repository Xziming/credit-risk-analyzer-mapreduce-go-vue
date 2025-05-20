package main

import (
    "credit/db"
    "credit/models"
    "log"
)

func main() {
    // 初始化数据库
    db.Init()
    db.DB.AutoMigrate(
        &models.AgeInfo{},
    )


    // 待插入数据（你可以继续补全完整数据） 
    ageDefaultsData := []models.AgeInfo{
        {Age: 21, Count: 67},
        {Age: 22, Count: 560},
        {Age: 23, Count: 931},
        {Age: 24, Count: 1127},
        {Age: 25, Count: 1186},
        {Age: 26, Count: 1256},
        {Age: 27, Count: 1477},
        {Age: 28, Count: 1409},
        {Age: 29, Count: 1605},
        {Age: 30, Count: 1395},
        {Age: 31, Count: 1217},
        {Age: 32, Count: 1158},
        {Age: 33, Count: 1146},
        {Age: 34, Count: 1162},
        {Age: 35, Count: 1113},
        {Age: 36, Count: 1108},
        {Age: 37, Count: 1041},
        {Age: 38, Count: 944},
        {Age: 39, Count: 954},
        {Age: 40, Count: 870},
        {Age: 41, Count: 824},
        {Age: 42, Count: 794},
        {Age: 43, Count: 670},
        {Age: 44, Count: 700},
        {Age: 45, Count: 617},
        {Age: 46, Count: 570},
        {Age: 47, Count: 501},
        {Age: 48, Count: 466},
        {Age: 49, Count: 452},
        {Age: 50, Count: 411},
        {Age: 51, Count: 340},
        {Age: 52, Count: 304},
        {Age: 53, Count: 325},
        {Age: 54, Count: 247},
        {Age: 55, Count: 209},
        {Age: 56, Count: 178},
        {Age: 57, Count: 122},
        {Age: 58, Count: 122},
        {Age: 59, Count: 83},
        {Age: 60, Count: 67},
        {Age: 61, Count: 56},
        {Age: 62, Count: 44},
        {Age: 63, Count: 31},
        {Age: 64, Count: 31},
        {Age: 65, Count: 24},
        {Age: 66, Count: 25},
        {Age: 67, Count: 16},
        {Age: 68, Count: 5},
        {Age: 69, Count: 15},
        {Age: 70, Count: 10},
        {Age: 71, Count: 3},
        {Age: 72, Count: 3},
        {Age: 73, Count: 4},
        {Age: 74, Count: 1},
        {Age: 75, Count: 3},
        {Age: 79, Count: 1},
    }


    // 插入数据
    for _, item := range ageData {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d 失败: %v\n", item.Age, item.Defaulted, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d 成功\n", item.Age, item.Defaulted)
        }
    }

    log.Println("数据插入完成。")
}

