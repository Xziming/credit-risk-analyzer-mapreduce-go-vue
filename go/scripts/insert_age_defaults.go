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
        &models.AgeDefaultsStats{},
    )


    // 待插入数据 
    data := []models.AgeDefaultsStats{
        {Age: 21, Defaulted: 0, Count: 53},
        {Age: 21, Defaulted: 1, Count: 14},
        {Age: 22, Defaulted: 0, Count: 391},
        {Age: 22, Defaulted: 1, Count: 169},
        {Age: 23, Defaulted: 0, Count: 684},
        {Age: 23, Defaulted: 1, Count: 247},
        {Age: 24, Defaulted: 0, Count: 827},
        {Age: 24, Defaulted: 1, Count: 300},
        {Age: 25, Defaulted: 0, Count: 884},
        {Age: 25, Defaulted: 1, Count: 302},
        {Age: 26, Defaulted: 0, Count: 1003},
        {Age: 26, Defaulted: 1, Count: 253},
        {Age: 27, Defaulted: 0, Count: 1164},
        {Age: 27, Defaulted: 1, Count: 313},
        {Age: 28, Defaulted: 0, Count: 1123},
        {Age: 28, Defaulted: 1, Count: 286},
        {Age: 29, Defaulted: 0, Count: 1292},
        {Age: 29, Defaulted: 1, Count: 313},
        {Age: 30, Defaulted: 0, Count: 1121},
        {Age: 30, Defaulted: 1, Count: 274},
        {Age: 31, Defaulted: 0, Count: 988},
        {Age: 31, Defaulted: 1, Count: 229},
        {Age: 32, Defaulted: 0, Count: 933},
        {Age: 32, Defaulted: 1, Count: 225},
        {Age: 33, Defaulted: 0, Count: 931},
        {Age: 33, Defaulted: 1, Count: 215},
        {Age: 34, Defaulted: 0, Count: 931},
        {Age: 34, Defaulted: 1, Count: 231},
        {Age: 35, Defaulted: 0, Count: 887},
        {Age: 35, Defaulted: 1, Count: 226},
        {Age: 36, Defaulted: 0, Count: 854},
        {Age: 36, Defaulted: 1, Count: 254},
        {Age: 37, Defaulted: 0, Count: 812},
        {Age: 37, Defaulted: 1, Count: 229},
        {Age: 38, Defaulted: 0, Count: 750},
        {Age: 38, Defaulted: 1, Count: 194},
        {Age: 39, Defaulted: 0, Count: 755},
        {Age: 39, Defaulted: 1, Count: 199},
        {Age: 40, Defaulted: 0, Count: 683},
        {Age: 40, Defaulted: 1, Count: 187},
        {Age: 41, Defaulted: 0, Count: 639},
        {Age: 41, Defaulted: 1, Count: 185},
        {Age: 42, Defaulted: 0, Count: 609},
        {Age: 42, Defaulted: 1, Count: 185},
        {Age: 43, Defaulted: 0, Count: 520},
        {Age: 43, Defaulted: 1, Count: 150},
        {Age: 44, Defaulted: 0, Count: 538},
        {Age: 44, Defaulted: 1, Count: 162},
        {Age: 45, Defaulted: 0, Count: 501},
        {Age: 45, Defaulted: 1, Count: 116},
        {Age: 46, Defaulted: 0, Count: 413},
        {Age: 46, Defaulted: 1, Count: 157},
        {Age: 47, Defaulted: 0, Count: 381},
        {Age: 47, Defaulted: 1, Count: 120},
        {Age: 48, Defaulted: 0, Count: 362},
        {Age: 48, Defaulted: 1, Count: 104},
        {Age: 49, Defaulted: 0, Count: 333},
        {Age: 49, Defaulted: 1, Count: 119},
        {Age: 50, Defaulted: 0, Count: 310},
        {Age: 50, Defaulted: 1, Count: 101},
        {Age: 51, Defaulted: 0, Count: 252},
        {Age: 51, Defaulted: 1, Count: 88},
        {Age: 52, Defaulted: 0, Count: 226},
        {Age: 52, Defaulted: 1, Count: 78},
        {Age: 53, Defaulted: 0, Count: 251},
        {Age: 53, Defaulted: 1, Count: 74},
        {Age: 54, Defaulted: 0, Count: 191},
        {Age: 54, Defaulted: 1, Count: 56},
        {Age: 55, Defaulted: 0, Count: 152},
        {Age: 55, Defaulted: 1, Count: 57},
        {Age: 56, Defaulted: 0, Count: 129},
        {Age: 56, Defaulted: 1, Count: 49},
        {Age: 57, Defaulted: 0, Count: 95},
        {Age: 57, Defaulted: 1, Count: 27},
        {Age: 58, Defaulted: 0, Count: 91},
        {Age: 58, Defaulted: 1, Count: 31},
        {Age: 59, Defaulted: 0, Count: 62},
        {Age: 59, Defaulted: 1, Count: 21},
        {Age: 60, Defaulted: 0, Count: 44},
        {Age: 60, Defaulted: 1, Count: 23},
        {Age: 61, Defaulted: 0, Count: 35},
        {Age: 61, Defaulted: 1, Count: 21},
        {Age: 62, Defaulted: 0, Count: 37},
        {Age: 62, Defaulted: 1, Count: 7},
        {Age: 63, Defaulted: 0, Count: 23},
        {Age: 63, Defaulted: 1, Count: 8},
        {Age: 64, Defaulted: 0, Count: 22},
        {Age: 64, Defaulted: 1, Count: 9},
        {Age: 65, Defaulted: 0, Count: 19},
        {Age: 65, Defaulted: 1, Count: 5},
        {Age: 66, Defaulted: 0, Count: 18},
        {Age: 66, Defaulted: 1, Count: 7},
        {Age: 67, Defaulted: 0, Count: 11},
        {Age: 67, Defaulted: 1, Count: 5},
        {Age: 68, Defaulted: 0, Count: 4},
        {Age: 68, Defaulted: 1, Count: 1},
        {Age: 69, Defaulted: 0, Count: 12},
        {Age: 69, Defaulted: 1, Count: 3},
        {Age: 70, Defaulted: 0, Count: 8},
        {Age: 70, Defaulted: 1, Count: 2},
        {Age: 71, Defaulted: 0, Count: 3},
        {Age: 72, Defaulted: 0, Count: 2},
        {Age: 72, Defaulted: 1, Count: 1},
        {Age: 73, Defaulted: 0, Count: 1},
        {Age: 73, Defaulted: 1, Count: 3},
        {Age: 74, Defaulted: 0, Count: 1},
        {Age: 75, Defaulted: 0, Count: 2},
        {Age: 75, Defaulted: 1, Count: 1},
        {Age: 79, Defaulted: 0, Count: 1},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d 失败: %v\n", item.Age, item.Defaulted, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d 成功\n", item.Age, item.Defaulted)
        }
    }

    log.Println("数据插入完成。")
}

