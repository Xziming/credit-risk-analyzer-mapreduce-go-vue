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
        &models.EducationDefaultsStats{},
    )


    // 待插入数据 
    data := []models.EducationDefaultsStats{
        {Education: 0, Defaulted: 0, Count: 14},
        {Education: 1, Defaulted: 0, Count: 8549},
        {Education: 1, Defaulted: 1, Count: 2036},
        {Education: 2, Defaulted: 0, Count: 10700},
        {Education: 2, Defaulted: 1, Count: 3330},
        {Education: 3, Defaulted: 0, Count: 3680},
        {Education: 3, Defaulted: 1, Count: 1237},
        {Education: 4, Defaulted: 0, Count: 116},
        {Education: 4, Defaulted: 1, Count: 7},
        {Education: 5, Defaulted: 0, Count: 262},
        {Education: 5, Defaulted: 1, Count: 18},
        {Education: 6, Defaulted: 0, Count: 43},
        {Education: 6, Defaulted: 1, Count: 8},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d, Count %d 失败: %v\n", item.Education, item.Defaulted, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d, Count %d 成功\n", item.Education, item.Defaulted, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

