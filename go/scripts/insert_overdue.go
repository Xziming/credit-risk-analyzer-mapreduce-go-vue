package main

import (
    "credit/db"
    "credit/models"
    "log"
)

func main() {
    db.Init()
    db.DB.AutoMigrate(
        &models.OverdueStreakInfo{},
    )

    data := []models.OverdueStreakInfo{
        {MaxOverdueStreak: 0, CustomerCount:19931},
        {MaxOverdueStreak: 1, CustomerCount:4842},
        {MaxOverdueStreak: 2, CustomerCount:1930},
        {MaxOverdueStreak: 3, CustomerCount:1013},
        {MaxOverdueStreak: 4, CustomerCount:657},
        {MaxOverdueStreak: 5, CustomerCount:286},
        {MaxOverdueStreak: 6, CustomerCount:1341},
    }

    for _, item := range data {
        if err := db.DB.Create(&item).Error; err != nil {
            log.Printf("Error: %d", err)
        }
    }
}
