package main

import (
    "credit/db"
    "credit/models"
    "log"
)

func main() {
    db.Init()
    db.DB.AutoMigrate(
        &models.BillAndRepayment{},
    )

    data := []models.BillAndRepayment{
        {Month:"4月", AverageBill: 39856.306, AverageRepay: 7314.43, OverdueAvgBill: 39154.236, NormalRepayRate: 0.1826},
        {Month: "5月", AverageBill: 41247.66, AverageRepay:6133.89 , OverdueAvgBill: 40367.459, NormalRepayRate: 0.1586},
        {Month: "6月", AverageBill: 44296.3191, AverageRepay: 6192.33, OverdueAvgBill: 42842.98, NormalRepayRate: 0.1462},
        {Month: "7月", AverageBill: 48095.248, AverageRepay: 6614.08, OverdueAvgBill: 46115.597, NormalRepayRate: 0.1429},
        {Month: "8月", AverageBill: 50328.85, AverageRepay: 7547.85, OverdueAvgBill: 48222.64, NormalRepayRate: 0.1541},
        {Month: "9月", AverageBill: 52274.099, AverageRepay: 7107.05, OverdueAvgBill: 49326.430, NormalRepayRate: 0.1387},
    }

    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("Error: %d", err)
        }
    }
}
