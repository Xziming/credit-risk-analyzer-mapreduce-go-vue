package routes

import (
    "credit/controllers"
    "credit/db"
    "credit/models"

    "github.com/gin-gonic/gin"

)

func Setup(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/", controllers.GetAll(db.DB, models.SexInfo{}))
        api.POST("/", controllers.Create[models.SexInfo](db.DB))
        api.PUT("/:id", controllers.Update(db.DB, models.SexInfo{}))
        api.DELETE("/:id", controllers.Delete(db.DB, models.SexInfo{}))
    }
}
