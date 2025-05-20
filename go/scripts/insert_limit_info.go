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
        &models.LimitBalInfo{},
    )


    // 待插入数据
    data := []models.LimitBalInfo{
        {Range: "10000", Count: 493},
        {Range: "100000", Count: 1048},
        {Range: "1000000", Count: 1},
        {Range: "110000", Count: 588},
        {Range: "120000", Count: 726},
        {Range: "130000", Count: 729},
        {Range: "140000", Count: 749},
        {Range: "150000", Count: 1110},
        {Range: "16000", Count: 2},
        {Range: "160000", Count: 694},
        {Range: "170000", Count: 532},
        {Range: "180000", Count: 995},
        {Range: "190000", Count: 229},
        {Range: "20000", Count: 1976},
        {Range: "200000", Count: 1528},
        {Range: "210000", Count: 730},
        {Range: "220000", Count: 469},
        {Range: "230000", Count: 737},
        {Range: "240000", Count: 619},
        {Range: "250000", Count: 350},
        {Range: "260000", Count: 521},
        {Range: "270000", Count: 238},
        {Range: "280000", Count: 493},
        {Range: "290000", Count: 348},
        {Range: "30000", Count: 1610},
        {Range: "300000", Count: 554},
        {Range: "310000", Count: 272},
        {Range: "320000", Count: 312},
        {Range: "327680", Count: 1},
        {Range: "330000", Count: 173},
        {Range: "340000", Count: 217},
        {Range: "350000", Count: 231},
        {Range: "360000", Count: 881},
        {Range: "370000", Count: 71},
        {Range: "380000", Count: 136},
        {Range: "390000", Count: 174},
        {Range: "40000", Count: 230},
        {Range: "400000", Count: 271},
        {Range: "410000", Count: 78},
        {Range: "420000", Count: 168},
        {Range: "430000", Count: 83},
        {Range: "440000", Count: 83},
        {Range: "450000", Count: 161},
        {Range: "460000", Count: 80},
        {Range: "470000", Count: 80},
        {Range: "480000", Count: 79},
        {Range: "490000", Count: 64},
        {Range: "50000", Count: 3365},
        {Range: "500000", Count: 722},
        {Range: "510000", Count: 19},
        {Range: "520000", Count: 20},
        {Range: "530000", Count: 10},
        {Range: "540000", Count: 6},
        {Range: "550000", Count: 21},
        {Range: "560000", Count: 10},
        {Range: "570000", Count: 8},
        {Range: "580000", Count: 11},
        {Range: "590000", Count: 6},
        {Range: "60000", Count: 825},
        {Range: "600000", Count: 16},
        {Range: "610000", Count: 11},
        {Range: "620000", Count: 9},
        {Range: "630000", Count: 7},
        {Range: "640000", Count: 7},
        {Range: "650000", Count: 3},
        {Range: "660000", Count: 3},
        {Range: "670000", Count: 3},
        {Range: "680000", Count: 4},
        {Range: "690000", Count: 1},
        {Range: "70000", Count: 731},
        {Range: "700000", Count: 8},
        {Range: "710000", Count: 6},
        {Range: "720000", Count: 3},
        {Range: "730000", Count: 2},
        {Range: "740000", Count: 2},
        {Range: "750000", Count: 4},
        {Range: "760000", Count: 1},
        {Range: "780000", Count: 2},
        {Range: "80000", Count: 1567},
        {Range: "800000", Count: 2},
        {Range: "90000", Count: 651},
    }


    // 插入数据
    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("插入 Age %d, Defaulted %d 失败: %v\n", item.Range, item.Count, err)
        } else {
            log.Printf("插入 Age %d, Defaulted %d 成功\n", item.Range, item.Count)
        }
    }

    log.Println("数据插入完成。")
}

