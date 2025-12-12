package handlers

import (
	"best-portfolio-go/database"
	"best-portfolio-go/models"
	"best-portfolio-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSkills handles GET /api/skills
func GetSkills(c *gin.Context) {
	var categories []models.SkillCategory
	db := database.GetDB()

	if err := db.Table("skill_categories").Order("created_at ASC").Find(&categories).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch skills: "+err.Error())
		return
	}

	utils.SuccessResponse(c, categories)
}

// GetSkillsPageData handles GET /api/skills/page-data
func GetSkillsPageData(c *gin.Context) {
	// This endpoint returns static metadata
	// In a real scenario, this could come from a database table
	pageData := models.SkillsPageData{
		Title:       "Skills",
		Description: "My technical skills and expertise",
	}

	utils.SuccessResponse(c, pageData)
}

// CreateSkillCategory handles POST /api/skills
func CreateSkillCategory(c *gin.Context) {
	var category models.SkillCategory

	if err := c.ShouldBindJSON(&category); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Ensure ID is not set (will be auto-generated)
	category.ID = 0

	db := database.GetDB()
	if err := db.Create(&category).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create skill category: "+err.Error())
		return
	}

	utils.SuccessResponse(c, category)
}

// UpdateSkillCategory handles PUT /api/skills/:id
func UpdateSkillCategory(c *gin.Context) {
	var category models.SkillCategory
	id := c.Param("id")

	categoryID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid skill category ID")
		return
	}

	db := database.GetDB()

	// Check if category exists
	if err := db.First(&category, categoryID).Error; err != nil {
		utils.NotFoundResponse(c, "Skill category")
		return
	}

	var updateData models.SkillCategory
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	if err := db.Model(&category).Updates(updateData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update skill category: "+err.Error())
		return
	}

	db.First(&category, categoryID)
	utils.SuccessResponse(c, category)
}

// DeleteSkillCategory handles DELETE /api/skills/:id
func DeleteSkillCategory(c *gin.Context) {
	var category models.SkillCategory
	id := c.Param("id")

	categoryID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid skill category ID")
		return
	}

	db := database.GetDB()

	// Check if category exists
	if err := db.First(&category, categoryID).Error; err != nil {
		utils.NotFoundResponse(c, "Skill category")
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete skill category: "+err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Skill category deleted successfully"})
}
