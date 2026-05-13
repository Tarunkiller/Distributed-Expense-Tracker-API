package controllers

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get monthly spending summary
// @Tags analytics
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/analytics/monthly [get]
func GetMonthlySummary(c *gin.Context) {
	userID, _ := c.Get("userID")

	var results []struct {
		Month string  `json:"month"`
		Total float64 `json:"total"`
	}

	// PostgreSQL specific date truncation
	err := config.DB.Model(&models.Expense{}).
		Select("TO_CHAR(date, 'YYYY-MM') as month, SUM(amount) as total").
		Where("user_id = ?", userID).
		Group("TO_CHAR(date, 'YYYY-MM')").
		Order("month DESC").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monthly summary"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// @Summary Get category-wise spending breakdown
// @Tags analytics
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/analytics/category [get]
func GetCategoryBreakdown(c *gin.Context) {
	userID, _ := c.Get("userID")

	var results []struct {
		Category string  `json:"category"`
		Total    float64 `json:"total"`
	}

	err := config.DB.Model(&models.Expense{}).
		Select("category, SUM(amount) as total").
		Where("user_id = ?", userID).
		Group("category").
		Order("total DESC").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category breakdown"})
		return
	}

	c.JSON(http.StatusOK, results)
}
