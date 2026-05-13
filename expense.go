package controllers

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseInput struct {
	Amount      float64   `json:"amount" binding:"required,gt=0"`
	Category    string    `json:"category" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
	TeamID      *uint     `json:"team_id,omitempty"`
}

// @Summary Create an expense
// @Tags expenses
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param expense body ExpenseInput true "Expense info"
// @Success 201 {object} models.Expense
// @Router /api/v1/expenses [post]
func CreateExpense(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input ExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := models.Expense{
		UserID:      userID.(uint),
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
		Date:        input.Date,
		TeamID:      input.TeamID,
	}

	if err := config.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

// @Summary Get all expenses
// @Tags expenses
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Param category query string false "Filter by category"
// @Success 200 {array} models.Expense
// @Router /api/v1/expenses [get]
func GetExpenses(c *gin.Context) {
	userID, _ := c.Get("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	category := c.Query("category")
	offset := (page - 1) * limit

	var expenses []models.Expense
	query := config.DB.Where("user_id = ?", userID)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	query.Offset(offset).Limit(limit).Order("date desc").Find(&expenses)

	c.JSON(http.StatusOK, expenses)
}

// @Summary Update an expense
// @Tags expenses
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Param expense body ExpenseInput true "Update info"
// @Success 200 {object} models.Expense
// @Router /api/v1/expenses/{id} [put]
func UpdateExpense(c *gin.Context) {
	userID, _ := c.Get("userID")
	expenseID := c.Param("id")

	var expense models.Expense
	if err := config.DB.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	var input ExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense.Amount = input.Amount
	expense.Category = input.Category
	expense.Description = input.Description
	expense.Date = input.Date

	config.DB.Save(&expense)
	c.JSON(http.StatusOK, expense)
}

// @Summary Delete an expense
// @Tags expenses
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/expenses/{id} [delete]
func DeleteExpense(c *gin.Context) {
	userID, _ := c.Get("userID")
	expenseID := c.Param("id")

	result := config.DB.Where("id = ? AND user_id = ?", expenseID, userID).Delete(&models.Expense{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}
