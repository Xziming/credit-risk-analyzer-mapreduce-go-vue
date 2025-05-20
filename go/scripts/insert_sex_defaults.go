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
        &models.GenderDefaultsStats{},
    )


    // 待插入数据 
    data := []models.GenderDefaultsStats{
        {Gender: 1, Defaulted: 0, Count: 9015},
        {Gender: 1, Defaulted: 1, Count: 2873},
        {Gender: 2, Defaulted: 0, Count: 14349},
        {Gender: 2, Defaulted: 1, Count: 3763},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d, Count %d 失败: %v\n", item.Gender, item.Defaulted, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d, Count %d 成功\n", item.Gender, item.Defaulted, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

