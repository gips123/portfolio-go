package handlers

import (
	"best-portfolio-go/database"
	"best-portfolio-go/models"
	"best-portfolio-go/utils"

	"github.com/gin-gonic/gin"
)

var validAboutIDs = map[string]bool{
	"about-me":     true,
	"aspirations":  true,
	"life-goals":   true,
	"hobbies":      true,
	"motto":        true,
}

// GetAboutCards handles GET /api/about
func GetAboutCards(c *gin.Context) {
	var cards []models.AboutCard
	db := database.GetDB()

	if err := db.Table("about_cards").Order("created_at ASC").Find(&cards).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch about cards: "+err.Error())
		return
	}

	utils.SuccessResponse(c, cards)
}

// GetAboutCardByID handles GET /api/about/:id
func GetAboutCardByID(c *gin.Context) {
	var card models.AboutCard
	id := c.Param("id")

	// Validate ID format
	if !validAboutIDs[id] {
		utils.BadRequestResponse(c, "Invalid about card ID. Valid IDs: about-me, aspirations, life-goals, hobbies, motto")
		return
	}

	db := database.GetDB()
	if err := db.First(&card, "id = ?", id).Error; err != nil {
		utils.NotFoundResponse(c, "About card")
		return
	}

	utils.SuccessResponse(c, card)
}

// GetAboutSidebar handles GET /api/about/sidebar
func GetAboutSidebar(c *gin.Context) {
	var cards []models.AboutCard
	db := database.GetDB()

	// Get all about cards to build sidebar
	if err := db.Order("created_at ASC").Find(&cards).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch about cards")
		return
	}

	// Convert to SidebarButton format
	sidebarButtons := make([]models.SidebarButton, len(cards))
	for i, card := range cards {
		sidebarButtons[i] = models.SidebarButton{
			ID:    card.ID,
			Title: card.Title,
			Icon:  card.Icon,
		}
	}

	utils.SuccessResponse(c, sidebarButtons)
}

// CreateAboutCard handles POST /api/about
func CreateAboutCard(c *gin.Context) {
	var card models.AboutCard

	if err := c.ShouldBindJSON(&card); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Validate ID
	if card.ID == "" {
		utils.BadRequestResponse(c, "ID is required")
		return
	}

	db := database.GetDB()

	// Check if ID already exists
	var existingCard models.AboutCard
	if err := db.First(&existingCard, "id = ?", card.ID).Error; err == nil {
		utils.BadRequestResponse(c, "About card with this ID already exists")
		return
	}

	if err := db.Create(&card).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create about card: "+err.Error())
		return
	}

	utils.SuccessResponse(c, card)
}

// UpdateAboutCard handles PUT /api/about/:id
func UpdateAboutCard(c *gin.Context) {
	var card models.AboutCard
	id := c.Param("id")

	db := database.GetDB()

	// Check if card exists
	if err := db.First(&card, "id = ?", id).Error; err != nil {
		utils.NotFoundResponse(c, "About card")
		return
	}

	var updateData models.AboutCard
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Don't allow changing ID
	updateData.ID = id

	if err := db.Model(&card).Updates(updateData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update about card: "+err.Error())
		return
	}

	db.First(&card, "id = ?", id)
	utils.SuccessResponse(c, card)
}

// DeleteAboutCard handles DELETE /api/about/:id
func DeleteAboutCard(c *gin.Context) {
	var card models.AboutCard
	id := c.Param("id")

	db := database.GetDB()

	// Check if card exists
	if err := db.First(&card, "id = ?", id).Error; err != nil {
		utils.NotFoundResponse(c, "About card")
		return
	}

	if err := db.Delete(&card).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete about card: "+err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "About card deleted successfully"})
}

