package controllers

import (
    "fmt"
    "net/http"
    "credit/utils"
    "credit/models"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.User
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
            return
        }

        var exist models.User
        if err := db.Where("username = ?", input.Username).First(&exist).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
            return
        }

        hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
            return
        }
        input.Password = string(hash)


        if err := db.Create(&input).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
    }
}

func Login(db *gorm.DB) gin.HandlerFunc {

    return func(c *gin.Context) {
        var input struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
            return
        }

        var user models.User
        if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名不存在"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
            return
        }
    token, err := utils.GenerateToken(user.Username)
    if err != nil {
        fmt.Printf("生成token失败: %v\n", err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
    }
}


