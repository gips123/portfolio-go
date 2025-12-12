package handlers

import (
	"best-portfolio-go/database"
	"best-portfolio-go/models"
	"best-portfolio-go/utils"

	"github.com/gin-gonic/gin"
)

// GetContact handles GET /api/contact
func GetContact(c *gin.Context) {
	var contactData models.ContactPageData
	db := database.GetDB()

	// Get the first (and typically only) contact data record
	// Use Select to explicitly specify columns
	if err := db.Table("contact_data").
		Select("id, title, description, contact_info, social_links, updated_at").
		First(&contactData).Error; err != nil {
		// If no data exists, return empty structure
		utils.SuccessResponse(c, models.ContactPageData{
			Title:       "Contact",
			Description: "Get in touch with me",
			ContactInfo: []models.ContactInfo{},
			SocialLinks: []models.SocialLink{},
		})
		return
	}

	utils.SuccessResponse(c, contactData)
}

// CreateOrUpdateContact handles POST /api/contact
func CreateOrUpdateContact(c *gin.Context) {
	var contactData models.ContactPageData
	db := database.GetDB()

	// Check if contact data already exists
	var existingData models.ContactPageData
	err := db.First(&existingData).Error

	if err != nil {
		// Create new
		if err := c.ShouldBindJSON(&contactData); err != nil {
			utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
			return
		}

		if err := db.Create(&contactData).Error; err != nil {
			utils.InternalServerErrorResponse(c, "Failed to create contact data: "+err.Error())
			return
		}

		utils.SuccessResponse(c, contactData)
	} else {
		// Update existing
		if err := c.ShouldBindJSON(&contactData); err != nil {
			utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
			return
		}

		if err := db.Model(&existingData).Updates(contactData).Error; err != nil {
			utils.InternalServerErrorResponse(c, "Failed to update contact data: "+err.Error())
			return
		}

		db.First(&existingData)
		utils.SuccessResponse(c, existingData)
	}
}

// UpdateContact handles PUT /api/contact
func UpdateContact(c *gin.Context) {
	var contactData models.ContactPageData
	db := database.GetDB()

	// Check if contact data exists
	var existingData models.ContactPageData
	if err := db.First(&existingData).Error; err != nil {
		utils.NotFoundResponse(c, "Contact data")
		return
	}

	if err := c.ShouldBindJSON(&contactData); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	if err := db.Model(&existingData).Updates(contactData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update contact data: "+err.Error())
		return
	}

	db.First(&existingData)
	utils.SuccessResponse(c, existingData)
}

// DeleteContact handles DELETE /api/contact
func DeleteContact(c *gin.Context) {
	var contactData models.ContactPageData
	db := database.GetDB()

	// Check if contact data exists
	if err := db.First(&contactData).Error; err != nil {
		utils.NotFoundResponse(c, "Contact data")
		return
	}

	if err := db.Delete(&contactData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete contact data: "+err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Contact data deleted successfully"})
}

