package routes

import (
	"expense-tracker/internal/controllers"
	"expense-tracker/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Auth Routes
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
	}

	// Protected API Routes
	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	{
		// Expenses
		api.POST("/expenses", controllers.CreateExpense)
		api.GET("/expenses", controllers.GetExpenses)
		api.PUT("/expenses/:id", controllers.UpdateExpense)
		api.DELETE("/expenses/:id", controllers.DeleteExpense)

		// Teams
		api.POST("/teams", controllers.CreateTeam)
		api.GET("/teams", controllers.GetTeams)

		// Analytics (placeholder, will be created in Task 5)
		api.GET("/analytics/monthly", controllers.GetMonthlySummary)
		api.GET("/analytics/category", controllers.GetCategoryBreakdown)
	}
}
