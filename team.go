package controllers

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// @Summary Create a team
// @Tags teams
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param team body TeamInput true "Team info"
// @Success 201 {object} models.Team
// @Router /api/v1/teams [post]
func CreateTeam(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input TeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	config.DB.First(&user, userID)

	team := models.Team{
		Name:        input.Name,
		Description: input.Description,
		CreatorID:   userID.(uint),
		Members:     []models.User{user}, // Creator is automatically a member
	}

	if err := config.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		return
	}

	c.JSON(http.StatusCreated, team)
}

// @Summary Get teams
// @Tags teams
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Team
// @Router /api/v1/teams [get]
func GetTeams(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	if err := config.DB.Preload("Teams").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teams"})
		return
	}

	// Wait, we need to setup Many2Many relationship properly if we want `user.Teams`.
	// For simplicity, let's query joining the association
	var teams []models.Team
	config.DB.Joins("JOIN team_members ON team_members.team_id = teams.id").
		Where("team_members.user_id = ?", userID).Find(&teams)

	c.JSON(http.StatusOK, teams)
}
