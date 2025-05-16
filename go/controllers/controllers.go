package controllers

import (
    "net/http"

    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "github.com/gin-gonic/gin"
)

// 通用获取
func GetAll[T any](db *gorm.DB, model T) gin.HandlerFunc {
    return func(c *gin.Context) {
        var data []T
        if err := db.Find(&data).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, data)
    }
}

// 通用创建
func Create[T any](db *gorm.DB, uniqueCols []string, updateCols []string ) gin.HandlerFunc {
    return func(c *gin.Context) {
        var entity T
        if err := c.ShouldBindJSON(&entity); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        cols := make([]clause.Column, len(uniqueCols))
        for i, col := range(uniqueCols) {
            cols[i] = clause.Column{Name: col}
        }

        assignments := clause.AssignmentColumns(updateCols)

        if err := db.Clauses(clause.OnConflict {
            Columns:    cols,
            DoUpdates:  assignments,
        }).Create(&entity).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, entity)
    }
}

// 通用更新
func Update[T any](db *gorm.DB, model T) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var input T
        if err := db.First(&input, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Save(&input).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, input)
    }
}

// 通用删除
func Delete[T any](db *gorm.DB, model T) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var entity T
        if err := db.Delete(&entity, id).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message":"Deleted"})
    }
}
