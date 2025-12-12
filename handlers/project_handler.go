package handlers

import (
	"best-portfolio-go/database"
	"best-portfolio-go/models"
	"best-portfolio-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var validCategories = map[string]bool{
	"all":           true,
	"frontend":      true,
	"backend":       true,
	"fullstack":     true,
	"uiux":          true,
	"mobile":        true,
	"desain-grafis": true,
}

// GetProjects handles GET /api/projects
func GetProjects(c *gin.Context) {
	var projects []models.Project
	db := database.GetDB()

	category := c.Query("category")

	// Validate category if provided
	if category != "" && !validCategories[category] {
		utils.BadRequestResponse(c, "Invalid category. Valid categories: all, frontend, backend, fullstack, uiux, mobile, desain-grafis")
		return
	}

	query := db.Table("projects")

	// Filter by category if provided and not "all"
	if category != "" && category != "all" {
		query = query.Where("category = ?", category)
	}

	// Order by updated_at DESC (newest first), then by created_at DESC as fallback
	if err := query.Order("updated_at DESC, created_at DESC").Find(&projects).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch projects: "+err.Error())
		return
	}

	utils.SuccessResponse(c, projects)
}

// GetProjectByID handles GET /api/projects/:id
func GetProjectByID(c *gin.Context) {
	var project models.Project
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	db := database.GetDB()
	if err := db.First(&project, projectID).Error; err != nil {
		utils.NotFoundResponse(c, "Project")
		return
	}

	utils.SuccessResponse(c, project)
}

// GetProjectCategories handles GET /api/projects/categories
func GetProjectCategories(c *gin.Context) {
	categories := []models.Category{
		{Value: "all", Label: "All"},
		{Value: "frontend", Label: "Frontend"},
		{Value: "backend", Label: "Backend"},
		{Value: "fullstack", Label: "Fullstack"},
		{Value: "uiux", Label: "UI/UX"},
		{Value: "mobile", Label: "Mobile"},
		{Value: "desain-grafis", Label: "Desain Grafis"},
	}

	utils.SuccessResponse(c, categories)
}

// GetProjectImages handles GET /api/projects/:id/images
func GetProjectImages(c *gin.Context) {
	var images []models.ProjectImage
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	// Verify project exists
	var project models.Project
	db := database.GetDB()
	if err := db.First(&project, projectID).Error; err != nil {
		utils.NotFoundResponse(c, "Project")
		return
	}

	// Get images ordered by order field (order is a reserved keyword in PostgreSQL)
	if err := db.Where("project_id = ?", projectID).
		Order(`"order" ASC`).
		Find(&images).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch project images")
		return
	}

	utils.SuccessResponse(c, images)
}

// CreateProject handles POST /api/projects
func CreateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Ensure ID is not set (will be auto-generated)
	project.ID = 0

	// Validate category
	if project.Category != "" && !validCategories[project.Category] {
		utils.BadRequestResponse(c, "Invalid category. Valid categories: frontend, backend, fullstack, uiux, mobile, desain-grafis")
		return
	}

	db := database.GetDB()
	if err := db.Create(&project).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create project: "+err.Error())
		return
	}

	utils.SuccessResponse(c, project)
}

// UpdateProject handles PUT /api/projects/:id
func UpdateProject(c *gin.Context) {
	var project models.Project
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	db := database.GetDB()

	// Check if project exists
	if err := db.First(&project, projectID).Error; err != nil {
		utils.NotFoundResponse(c, "Project")
		return
	}

	// Bind update data
	var updateData models.Project
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Validate category if provided
	if updateData.Category != "" && !validCategories[updateData.Category] {
		utils.BadRequestResponse(c, "Invalid category. Valid categories: frontend, backend, fullstack, uiux, mobile, desain-grafis")
		return
	}

	// Update project
	if err := db.Model(&project).Updates(updateData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update project: "+err.Error())
		return
	}

	// Reload to get updated data
	db.First(&project, projectID)
	utils.SuccessResponse(c, project)
}

// DeleteProject handles DELETE /api/projects/:id
func DeleteProject(c *gin.Context) {
	var project models.Project
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	db := database.GetDB()

	// Check if project exists
	if err := db.First(&project, projectID).Error; err != nil {
		utils.NotFoundResponse(c, "Project")
		return
	}

	// Delete project (cascade will delete related images)
	if err := db.Delete(&project).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete project: "+err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Project deleted successfully"})
}

// CreateProjectImage handles POST /api/projects/:id/images
func CreateProjectImage(c *gin.Context) {
	var image models.ProjectImage
	id := c.Param("id")

	projectID, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	// Verify project exists
	var project models.Project
	db := database.GetDB()
	if err := db.First(&project, projectID).Error; err != nil {
		utils.NotFoundResponse(c, "Project")
		return
	}

	if err := c.ShouldBindJSON(&image); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Ensure ID is not set (will be auto-generated)
	image.ID = 0
	image.ProjectID = projectID

	if err := db.Create(&image).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create project image: "+err.Error())
		return
	}

	utils.SuccessResponse(c, image)
}

// UpdateProjectImage handles PUT /api/projects/:id/images/:imageId
func UpdateProjectImage(c *gin.Context) {
	var image models.ProjectImage
	imageID := c.Param("imageId")
	projectID := c.Param("id")

	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	imageIDInt, err := strconv.Atoi(imageID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid image ID")
		return
	}

	db := database.GetDB()

	// Check if image exists and belongs to project
	if err := db.Where("id = ? AND project_id = ?", imageIDInt, projectIDInt).First(&image).Error; err != nil {
		utils.NotFoundResponse(c, "Project image")
		return
	}

	var updateData models.ProjectImage
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	// Don't allow changing project_id
	updateData.ProjectID = projectIDInt

	if err := db.Model(&image).Updates(updateData).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update project image: "+err.Error())
		return
	}

	db.First(&image, imageIDInt)
	utils.SuccessResponse(c, image)
}

// DeleteProjectImage handles DELETE /api/projects/:id/images/:imageId
func DeleteProjectImage(c *gin.Context) {
	var image models.ProjectImage
	imageID := c.Param("imageId")
	projectID := c.Param("id")

	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid project ID")
		return
	}

	imageIDInt, err := strconv.Atoi(imageID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid image ID")
		return
	}

	db := database.GetDB()

	// Check if image exists and belongs to project
	if err := db.Where("id = ? AND project_id = ?", imageIDInt, projectIDInt).First(&image).Error; err != nil {
		utils.NotFoundResponse(c, "Project image")
		return
	}

	if err := db.Delete(&image).Error; err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete project image: "+err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Project image deleted successfully"})
}
