package routes

import (
    "credit/controllers"
    "credit/db"
    "credit/models"

    "github.com/gin-gonic/gin"

)

func Setup(r *gin.Engine) {
    db := db.DB
    api := r.Group("/api")
    {
        // User
        api.POST("/register", controllers.Register(db))
        api.POST("/login", controllers.Login(db))
    }
    {
        // 基本客户信息
        api.GET("/clientsinfo", controllers.GetAll[models.ClientsInfo](db, models.ClientsInfo{}))

    }
    {
        // 性别分布
        api.GET("/genderinfo", controllers.GetAll[models.GenderInfo](db, models.GenderInfo{}))
        api.POST("/genderinfo", controllers.Create[models.GenderInfo](db, []string{"gender"}, []string{"count"}))
        api.PUT("/genderinfo/:id", controllers.Update[models.GenderInfo](db, models.GenderInfo{}))
        api.DELETE("/genderinfo/:id", controllers.Delete[models.GenderInfo](db, models.GenderInfo{}))
    }

    {
        // 年龄分布
        api.GET("/ageinfo", controllers.GetAll[models.AgeInfo](db, models.AgeInfo{}))
        api.POST("/ageinfo", controllers.Create[models.AgeInfo](db, []string{"age"}, []string{"count"}))
        api.PUT("/ageinfo/:id", controllers.Update[models.AgeInfo](db, models.AgeInfo{}))
        api.DELETE("/ageinfo/:id", controllers.Delete[models.AgeInfo](db, models.AgeInfo{}))
    }

    {
        // 教育分布
        api.GET("/educationinfo", controllers.GetAll[models.EducationInfo](db, models.EducationInfo{}))
        api.POST("/educationinfo", controllers.Create[models.EducationInfo](db, []string{"education"}, []string{"count"}))
        api.PUT("/educationinfo/:id", controllers.Update[models.EducationInfo](db, models.EducationInfo{}))
        api.DELETE("/educationinfo/:id", controllers.Delete[models.EducationInfo](db, models.EducationInfo{}))
    }

    {
        // 婚姻分布
        api.GET("/marriageinfo", controllers.GetAll[models.MarriageInfo](db, models.MarriageInfo{}))
        api.POST("/marriageinfo", controllers.Create[models.MarriageInfo](db, []string{"marriage"}, []string{"count"}))
        api.PUT("/marriageinfo/:id", controllers.Update[models.MarriageInfo](db, models.MarriageInfo{}))
        api.DELETE("/marriageinfo/:id", controllers.Delete[models.MarriageInfo](db, models.MarriageInfo{}))
    }

    {
        // 信用额度范围分布
        api.GET("/limitbalinfo", controllers.GetAll[models.LimitBalInfo](db, models.LimitBalInfo{}))
        api.POST("/limitbalinfo", controllers.Create[models.LimitBalInfo](db, []string{"range"}, []string{"count"}))
        api.PUT("/limitbalinfo/:id", controllers.Update[models.LimitBalInfo](db, models.LimitBalInfo{}))
        api.DELETE("/limitbalinfo/:id", controllers.Delete[models.LimitBalInfo](db, models.LimitBalInfo{}))
    }

    {
        // 按性别违约统计
        api.GET("/genderdefaults", controllers.GetAll[models.GenderDefaultsStats](db, models.GenderDefaultsStats{}))
        api.POST("/genderdefaults", controllers.Create[models.GenderDefaultsStats](db, []string{"gender", "defaulted"}, []string{"count"}))
        api.PUT("/genderdefaults/:id", controllers.Update[models.GenderDefaultsStats](db, models.GenderDefaultsStats{}))
        api.DELETE("/genderdefaults/:id", controllers.Delete[models.GenderDefaultsStats](db, models.GenderDefaultsStats{}))
    }

    {
        // 按年龄违约统计
        api.GET("/agedefaults", controllers.GetAll[models.AgeDefaultsStats](db, models.AgeDefaultsStats{}))
        api.POST("/agedefaults", controllers.Create[models.AgeDefaultsStats](db, []string{"age", "defaulted"}, []string{"count"}))
        api.PUT("/agedefaults/:id", controllers.Update[models.AgeDefaultsStats](db, models.AgeDefaultsStats{}))
        api.DELETE("/agedefaults/:id", controllers.Delete[models.AgeDefaultsStats](db, models.AgeDefaultsStats{}))
    }

    {
        // 按教育违约统计
        api.GET("/educationdefaults", controllers.GetAll[models.EducationDefaultsStats](db, models.EducationDefaultsStats{}))
        api.POST("/educationdefaults", controllers.Create[models.EducationDefaultsStats](db, []string{"education", "defaulted"}, []string{"count"}))
        api.PUT("/educationdefaults/:id", controllers.Update[models.EducationDefaultsStats](db, models.EducationDefaultsStats{}))
        api.DELETE("/educationdefaults/:id", controllers.Delete[models.EducationDefaultsStats](db, models.EducationDefaultsStats{}))
    }

    {
        // 按婚姻违约统计
        api.GET("/marriagedefaults", controllers.GetAll[models.MarriageDefaultsStats](db, models.MarriageDefaultsStats{}))
        api.POST("/marriagedefaults", controllers.Create[models.MarriageDefaultsStats](db, []string{"marriage", "defaulted"}, []string{"count"}))
        api.PUT("/marriagedefaults/:id", controllers.Update[models.MarriageDefaultsStats](db, models.MarriageDefaultsStats{}))
        api.DELETE("/marriagedefaults/:id", controllers.Delete[models.MarriageDefaultsStats](db, models.MarriageDefaultsStats{}))
    }

    {
        // 按额度违约统计
        api.GET("/limitdefaults", controllers.GetAll[models.LimitDefaultsStats](db, models.LimitDefaultsStats{}))
        api.POST("/limitdefaults", controllers.Create[models.LimitDefaultsStats](db, []string{"range", "defaulted"}, []string{"count"}))
        api.PUT("/limitdefaults/:id", controllers.Update[models.LimitDefaultsStats](db, models.LimitDefaultsStats{}))
        api.DELETE("/limitdefaults/:id", controllers.Delete[models.LimitDefaultsStats](db, models.LimitDefaultsStats{}))
    }

    {
        // 账单还款对比
        api.GET("/billpaycomparison", controllers.GetAll[models.BillPayComparison](db, models.BillPayComparison{}))
        api.POST("/billpaycomparison", controllers.Create[models.BillPayComparison](db, []string{"month"}, []string{"total_bill", "total_pay"}))
        api.PUT("/billpaycomparison/:id", controllers.Update[models.BillPayComparison](db, models.BillPayComparison{}))
        api.DELETE("/billpaycomparison/:id", controllers.Delete[models.BillPayComparison](db, models.BillPayComparison{}))
    }

    {
        // 连续违约统计
        api.GET("/consecutivedefaults", controllers.GetAll[models.ConsecutiveDefaults](db, models.ConsecutiveDefaults{}))
        api.POST("/consecutivedefaults", controllers.Create[models.ConsecutiveDefaults](db, []string{"client_id"}, []string{"max_consecutive"}))
        api.PUT("/consecutivedefaults/:id", controllers.Update[models.ConsecutiveDefaults](db, models.ConsecutiveDefaults{}))
        api.DELETE("/consecutivedefaults/:id", controllers.Delete[models.ConsecutiveDefaults](db, models.ConsecutiveDefaults{}))
    }

    {
        // 综合评分汇总
        api.GET("/defaultscoresummary", controllers.GetAll[models.DefaultScoreSummary](db, models.DefaultScoreSummary{}))
        api.POST("/defaultscoresummary", controllers.Create[models.DefaultScoreSummary](db, []string{"score_name"}, []string{"client_count"}))
        api.PUT("/defaultscoresummary/:id", controllers.Update[models.DefaultScoreSummary](db, models.DefaultScoreSummary{}))
        api.DELETE("/defaultscoresummary/:id", controllers.Delete[models.DefaultScoreSummary](db, models.DefaultScoreSummary{}))
    }

    {
        api.GET("/billandrepayment", controllers.GetAll[models.BillAndRepayment](db, models.BillAndRepayment{}))
        api.GET("/overdue", controllers.GetAll[models.OverdueStreakInfo](db, models.OverdueStreakInfo{}))
    }
}
