package db

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库连接 读取 .env 中的配置
func Init() {
    // 加载 .env 文件
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    } else {
        log.Println(".env file loaded successfully")
    }

    // 构造 DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
    os.Getenv("DB_NAME"),
)
log.Println(dsn)

// 连接数据库
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
    log.Fatal("Failed to connect database:", err)
}
DB = db
log.Println("Datebase connection established")
}
