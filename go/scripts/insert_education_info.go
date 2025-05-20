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
        &models.EducationInfo{},
    )


    // 待插入数据
    data := []models.EducationInfo{
        {Education: 0, Count: 14},
        {Education: 1, Count: 10585},
        {Education: 2, Count: 14030},
        {Education: 3, Count: 4917},
        {Education: 4, Count: 123},
        {Education: 5, Count: 280},
        {Education: 6, Count: 51},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d 失败: %v\n", item.Education, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d 成功\n", item.Education, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

