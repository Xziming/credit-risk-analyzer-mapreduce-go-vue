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
        &models.MarriageInfo{},
    )


    // 待插入数据
    ageData := []models.MarriageInfo{
        {Marriage: 0, Count: 54},
        {Marriage: 1, Count: 13659},
        {Marriage: 2, Count: 15964},
        {Marriage: 3, Count: 323},
    }


    // 插入数据
    for _, item := range ageData {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d 失败: %v\n", item.Marriage, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d 成功\n", item.Marriage, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

