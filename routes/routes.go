package routes

import (
	"best-portfolio-go/config"
	"best-portfolio-go/handlers"
	"best-portfolio-go/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     cfg.CORS.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	// API routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.LoginHandler(cfg))
		}

		// Projects routes
		projects := api.Group("/projects")
		{
			projects.GET("", handlers.GetProjects)
			projects.GET("/categories", handlers.GetProjectCategories)
			projects.GET("/:id", handlers.GetProjectByID)
			projects.GET("/:id/images", handlers.GetProjectImages)
		}
		protectedProjects := api.Group("/projects", middleware.RequireAuth(cfg))
		{
			protectedProjects.POST("", handlers.CreateProject)
			protectedProjects.PUT("/:id", handlers.UpdateProject)
			protectedProjects.DELETE("/:id", handlers.DeleteProject)
			protectedProjects.POST("/:id/images", handlers.CreateProjectImage)
			protectedProjects.PUT("/:id/images/:imageId", handlers.UpdateProjectImage)
			protectedProjects.DELETE("/:id/images/:imageId", handlers.DeleteProjectImage)
		}

		// About routes
		about := api.Group("/about")
		{
			about.GET("", handlers.GetAboutCards)
			about.GET("/sidebar", handlers.GetAboutSidebar)
			about.GET("/:id", handlers.GetAboutCardByID)
		}
		protectedAbout := api.Group("/about", middleware.RequireAuth(cfg))
		{
			protectedAbout.POST("", handlers.CreateAboutCard)
			protectedAbout.PUT("/:id", handlers.UpdateAboutCard)
			protectedAbout.DELETE("/:id", handlers.DeleteAboutCard)
		}

		// Skills routes
		skills := api.Group("/skills")
		{
			skills.GET("", handlers.GetSkills)
			skills.GET("/page-data", handlers.GetSkillsPageData)
		}
		protectedSkills := api.Group("/skills", middleware.RequireAuth(cfg))
		{
			protectedSkills.POST("", handlers.CreateSkillCategory)
			protectedSkills.PUT("/:id", handlers.UpdateSkillCategory)
			protectedSkills.DELETE("/:id", handlers.DeleteSkillCategory)
		}

		// Contact routes
		contact := api.Group("/contact")
		{
			contact.GET("", handlers.GetContact)
		}
		protectedContact := api.Group("/contact", middleware.RequireAuth(cfg))
		{
			protectedContact.POST("", handlers.CreateOrUpdateContact)
			protectedContact.PUT("", handlers.UpdateContact)
			protectedContact.DELETE("", handlers.DeleteContact)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Portfolio API is running",
		})
	})

	return router
}
