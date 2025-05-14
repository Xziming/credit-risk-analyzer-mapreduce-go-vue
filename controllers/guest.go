package controllers

import (
    "net/http"

    "gorm-v1/database"
    "gorm-v1/models"

    "github.com/gin-gonic/gin"
)

func AddNewGuest(c *gin.Context){
    var guest models.Guest
    if err := c.ShouldBindJSON(&guest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
     }

    if err := database.DB.Create(&guest).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "success", "guest": guest.ID})
}

func GetGuestsInfo(c *gin.Context) {


