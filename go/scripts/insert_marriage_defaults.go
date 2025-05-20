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
        &models.MarriageDefaultsStats{},
    )

    // 数据
    data := []models.MarriageDefaultsStats{
        {Marriage: 0, Defaulted: 0, Count: 49},
        {Marriage: 0, Defaulted: 1, Count: 5},
        {Marriage: 1, Defaulted: 0, Count: 10453},
        {Marriage: 1, Defaulted: 1, Count: 3206},
        {Marriage: 2, Defaulted: 0, Count: 12623},
        {Marriage: 2, Defaulted: 1, Count: 3341},
        {Marriage: 3, Defaulted: 0, Count: 239},
        {Marriage: 3, Defaulted: 1, Count: 84},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d, Count %d 失败: %v\n", item.Marriage, item.Defaulted, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d, Count %d 成功\n", item.Marriage, item.Defaulted, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

