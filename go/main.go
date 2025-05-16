package main

import (
    "log"

    "credit/db"
    "credit/models"
    "credit/routes"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // 加载 .env 配置
    if err := godotenv.Load(); err != nil {
        log.Println("Error, No .env file found")
    }

    // 初始化数据库，自动迁移 models 到数据库
    db.Init()
    db.DB.AutoMigrate(
        &models.ClientsInfo{},
        &models.GenderInfo{},
        &models.AgeInfo{},
        &models.EducationInfo{},
        &models.MarriageInfo{},
        &models.LimitBalInfo{},
        &models.GenderDefaultsStats{},
        &models.AgeDefaultsStats{},
        &models.EducationDefaultsStats{},
        &models.MarriageDefaultsStats{},
        &models.LimitDefaultsStats{},
        &models.BillPayComparison{},
        &models.ConsecutiveDefaults{},
        &models.DefaultScoreSummary{},
    )

    r := gin.Default()


    routes.Setup(r)

    // 启动服务器
    log.Println("Server running at 8.155.56.96:9001")
    if err := r.Run(":9001"); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}
