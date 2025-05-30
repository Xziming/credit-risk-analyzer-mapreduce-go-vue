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
        &models.GenderInfo{},
    )


    // 待插入数据
    data := []models.GenderInfo{
        {Gender: 1, Count: 11888},
        {Gender: 2, Count: 18112},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d 失败: %v\n", item.Gender, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d 成功\n", item.Gender, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

